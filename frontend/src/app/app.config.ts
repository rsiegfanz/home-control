import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { ApplicationConfig, inject } from '@angular/core';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { provideRouter } from '@angular/router';
import { routes } from './app.routes';
import { provideApollo } from 'apollo-angular';
import { HttpLink } from 'apollo-angular/http';
import { InMemoryCache, split } from '@apollo/client/core';
import { environment } from './environments/environment';
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import { getMainDefinition } from '@apollo/client/utilities';
import { createClient } from 'graphql-ws';

export const appConfig: ApplicationConfig = {
    providers: [
        provideRouter(routes),
        provideAnimationsAsync(),
        provideHttpClient(),
        provideApollo(() => {
            const httpLink = inject(HttpLink);

            const httpUrl = `http://${environment.backendGoUrl}/graphql`; // todo http
            const http = httpLink.create({
                uri: httpUrl,
            });

            const wsUrl = `ws://${environment.backendGoUrl}/graphql`;
            const wsClient = createClient({
                url: wsUrl,
                connectionParams: {},
                on: {
                    connected: () => console.log('WebSocket connected'),
                    error: (error) => console.error('WebSocket error:', error),
                    closed: () => console.log('WebSocket connection closed'),
                },
            });
            const ws = new GraphQLWsLink(wsClient);

            ws.client.on('connected', () => console.log('connected'));
            ws.client.on('error', (error) => console.log('error', error));

            const link = split(
                ({ query }) => {
                    const definition = getMainDefinition(query);
                    return definition.kind === 'OperationDefinition' && definition.operation === 'subscription';
                },
                ws,
                http,
            );

            return {
                link,
                cache: new InMemoryCache(),
            };
        }),
    ],
};

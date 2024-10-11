import { provideHttpClient } from '@angular/common/http';
import { ApplicationConfig, inject } from '@angular/core';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { provideRouter } from '@angular/router';
import { routes } from './app.routes';
import { provideApollo } from 'apollo-angular';
import { HttpLink } from 'apollo-angular/http';
import { InMemoryCache } from '@apollo/client/core';
import { environment } from './environments/environment';

export const appConfig: ApplicationConfig = {
    providers: [
        provideRouter(routes),
        provideAnimationsAsync(),
        provideHttpClient(),
        provideApollo(() => {
            const httpLink = inject(HttpLink);

            return {
                link: httpLink.create({
                    uri: `${environment.backendGoUrl}/graphql`,
                }),
                cache: new InMemoryCache(),
            };
        }),
    ],
};

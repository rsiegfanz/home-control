import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { inject } from '@angular/core';
import { Observable, catchError, map, of } from 'rxjs';
import urlJoin from 'url-join';
import { environment } from '../../../environments/environment';
import { ApiResponse } from '../models/api-response.model';
import { BaseModel } from '../models/base.model';
import { IApiResponseDto } from '..//dtos/api-response-dto.interface';

export abstract class BackendService<TModel extends BaseModel> {
    protected http = inject(HttpClient);

    protected baseUrl = `http://${environment.backendGoUrl}`; // todo

    protected abstract subdirectory: string;

    protected getSingle<TDto>(url: string, mapper: (dto: TDto | undefined) => TModel | undefined): Observable<ApiResponse<TModel>> {
        return this.http.get<IApiResponseDto<TDto>>(url).pipe(
            map((responseDto: IApiResponseDto<TDto>) => this._responseHandling(responseDto, mapper)),
            catchError((error: HttpErrorResponse) => this._httpErrorToApiResponse(error)),
        );
    }

    protected getMultiple<TDto>(url: string, mapper: (dtos: TDto[] | undefined) => TModel[] | undefined): Observable<ApiResponse<TModel[]>> {
        return this.http.get<IApiResponseDto<TDto[]>>(url).pipe(
            map((responseDto: IApiResponseDto<TDto[]>) => this._responseHandlingMultiple(responseDto, mapper)),
            catchError((error: HttpErrorResponse) => this._httpErrorToApiResponseMultiple(error)),
        );
    }

    protected createUrl(path?: string): string {
        if (!path) {
            return urlJoin(this.baseUrl, this.subdirectory);
        }

        return urlJoin(this.baseUrl, this.subdirectory, path);
    }

    private _responseHandling<TDto>(responseDto?: IApiResponseDto<TDto>, mapper?: (data?: TDto | undefined) => TModel | undefined): ApiResponse<TModel> {
        if (!responseDto) {
            return this._errorHandling(1000, ['no response delivered']);
        }

        if (!responseDto.status) {
            return this._errorHandling(1000, ['no status delivered']);
        }

        if (responseDto.status >= 400) {
            return this._errorHandling(responseDto.status, responseDto.messages);
        }

        if (!mapper) {
            return this._errorHandling(1000, ['no mapper provided']);
        }

        return ApiResponse.createSuccess(mapper(responseDto.data as TDto));
    }

    private _responseHandlingMultiple<TDto>(responseDto?: IApiResponseDto<TDto[]>, mapper?: (data?: [] | undefined) => TModel[] | undefined): ApiResponse<TModel[]> {
        if (!responseDto) {
            return this._errorHandlingMultiple(1000, ['no response delivered']);
        }

        if (!responseDto.status) {
            return this._errorHandlingMultiple(1000, ['no status delivered']);
        }

        if (responseDto.status >= 400) {
            return this._errorHandlingMultiple(responseDto.status, responseDto.messages);
        }

        if (!mapper) {
            return this._errorHandlingMultiple(1000, ['no mapper provided']);
        }

        return ApiResponse.createSuccessArray(mapper(responseDto.data as []));
    }

    private _errorHandling(statusCode: number, messages: string[] | undefined): ApiResponse<TModel> {
        console.log(`Error`);
        return ApiResponse.createError<TModel>(messages);
    }

    private _errorHandlingMultiple(statusCode: number, messages: string[] | undefined): ApiResponse<TModel[]> {
        console.log(`Error`);
        return ApiResponse.createErrorArray<TModel>(messages);
    }

    private _httpErrorToApiResponse(error: HttpErrorResponse): Observable<ApiResponse<TModel>> {
        const messages: string[] = [];
        if (error.message) {
            messages.push(error.message);
        }
        if (error.error?.messages) {
            messages.push(...error.error.messages);
        }
        return of(this._errorHandling(error.status, messages));
    }

    private _httpErrorToApiResponseMultiple(error: HttpErrorResponse): Observable<ApiResponse<TModel[]>> {
        const messages: string[] = [];
        if (error.message) {
            messages.push(error.message);
        }
        if (error.error?.messages) {
            messages.push(...error.error.messages);
        }
        return of(this._errorHandlingMultiple(error.status, messages));
    }
}

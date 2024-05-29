import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, map, of } from 'rxjs';
import urlJoin from 'url-join';
import { IApiODataResponseDto } from '../../dtos/api-odata-response-dto.interface';
import { IApiResponseDto } from '../../dtos/api-response-dto.interface';
import { IBaseRequestDto } from '../../dtos/base-request-dto.interface';
import { IBaseResponseDto } from '../../dtos/base-response-dto.interface';
import { ApiResponse } from '../../models/api-response.model';
import { BaseModel } from '../../models/base.model';
import { ODataPagination } from '../../models/odata/odata-pagination.model';
import { OdataResponse } from '../../models/odata/odata-response.interface';
import { OData } from '../../models/odata/odata.model';

@Injectable({
    providedIn: 'root',
})
export abstract class BaseApiRepository<TModel extends BaseModel, TResponseDto extends IBaseResponseDto, TRequestDto extends IBaseRequestDto> {
    protected abstract readonly url: string;

    protected abstract readonly path: string;

    public abstract mapResponseDtoToModel(dto: TResponseDto): TModel;
    public abstract mapModelToRequestDto(model: TModel): TRequestDto;

    constructor(protected readonly http: HttpClient) {}

    public get(path: string): Observable<ApiResponse<TModel>> {
        return this.http.get<IApiResponseDto<TResponseDto>>(path).pipe(
            map((response: IApiResponseDto<TResponseDto>) => this._responseMapping(response)),
            catchError((error: HttpErrorResponse) => {
                return this._httpErrorToApiResponse(error);
            }),
        );
    }

    public getById(id: string): Observable<ApiResponse<TModel>> {
        return this.http.get<IApiResponseDto<TResponseDto>>(`${this.url}/${id}`).pipe(
            map((response: IApiResponseDto<TResponseDto>) => this._responseMapping(response)),
            catchError((error: HttpErrorResponse) => {
                return this._httpErrorToApiResponse(error);
            }),
        );
    }

    public getAll(): Observable<ApiResponse<TModel>> {
        const odata = new OData();
        return this.odata(odata);
    }

    public odata(odataQuery: OData): Observable<ApiResponse<TModel>> {
        return this.http.get<IApiODataResponseDto<TResponseDto>>(`${this.url}/${odataQuery.toGetParameter()}`).pipe(
            map((response: IApiODataResponseDto<TResponseDto>) => this._responseOdataMapping(response)),
            catchError((error: HttpErrorResponse) => {
                return this._httpErrorToApiResponse(error);
            }),
        );
    }

    public insert(model: TModel): Observable<ApiResponse<TModel>> {
        const dto = this.mapModelToRequestDto(model);
        return this.http.post<IApiResponseDto<TResponseDto>>(this.url, dto).pipe(
            map((response: IApiResponseDto<TResponseDto>) => this._responseMapping(response)),
            catchError((error: HttpErrorResponse) => {
                return this._httpErrorToApiResponse(error);
            }),
        );
    }

    private _responseOdataMapping(response: IApiODataResponseDto<TResponseDto>): ApiResponse<TModel> {
        if (response.status !== 200) {
            return this._errorHandling(response.status, ['query failed']);
        }
        const apiResponse = new ApiResponse<TModel>();
        apiResponse.status = response.status;
        const odataResponse = new OdataResponse<TModel>();
        odataResponse.items = response.data.items.map((dto) => this.mapResponseDtoToModel(dto));
        odataResponse.pagination = ODataPagination.fromDto(response.data.pagination);
        apiResponse.odata = odataResponse;
        return apiResponse;
    }

    private _responseMapping(response: IApiResponseDto<TResponseDto>): ApiResponse<TModel> {
        if (![200, 201].includes(response.status) || !response.data) {
            return this._errorHandling(response.status, response.messages);
        }
        const apiResponse = new ApiResponse<TModel>();
        apiResponse.status = response.status;
        apiResponse.messages = response.messages;
        apiResponse.isError = false;

        apiResponse.data = this.mapResponseDtoToModel(response.data);
        return apiResponse;
    }

    private _errorHandling(statusCode: number, messages: string[]): ApiResponse<TModel> {
        const apiResponse = new ApiResponse<TModel>();
        apiResponse.status = statusCode;
        apiResponse.messages = messages;
        apiResponse.isError = true;
        return apiResponse;
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

    protected urlCombine(): string {
        return urlJoin(this.url, this.path);
    }
}

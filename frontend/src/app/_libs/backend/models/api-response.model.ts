import { BaseModel } from './base.model';

export class ApiResponse<TModel extends BaseModel> {
    public readonly data: TModel | undefined;

    public readonly messages: string[] | undefined;

    public readonly isSuccessful!: boolean;

    private constructor(isSuccessful: boolean, data?: TModel | undefined, messages?: string[] | undefined) {
        this.isSuccessful = isSuccessful;
        this.data = data;
        this.messages = messages;
    }

    public static createSuccess<TModel extends BaseModel>(data?: TModel | undefined): ApiResponse<TModel> {
        return new ApiResponse<TModel>(true, data);
    }

    public static createSuccessArray<TModel extends BaseModel>(data?: TModel[] | undefined): ApiResponse<TModel[]> {
        return new ApiResponse<TModel[]>(true, data);
    }

    public static createError<TModel extends BaseModel>(messages?: string[] | undefined): ApiResponse<TModel> {
        return new ApiResponse<TModel>(false, undefined, messages);
    }

    public static createErrorArray<TModel extends BaseModel>(messages?: string[] | undefined): ApiResponse<TModel[]> {
        return new ApiResponse<TModel[]>(false, undefined, messages);
    }
}

export interface IApiResponseDto<TDto> {
    messages: string[] | undefined;
    data: TDto | TDto[] | undefined;
    status: number | undefined;
}

import axios, { AxiosResponse } from 'axios';
import { ListItems as ListItemsPath } from '@/services/endpoint';
import { IApiReturn, IListItems } from '@/store/typings/root';

export const ListProductService = () => {
  return axios({
    url: ListItemsPath,
    method: 'GET'
  }) as Promise<AxiosResponse<IApiReturn<IListItems[]>, any>>;
};

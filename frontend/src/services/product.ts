import axios, { AxiosResponse } from 'axios';
import { ListItems as ListItemsPath, BuyItems } from '@/services/endpoint';
import { IApiReturn, IListItems } from '@/store/typings/root';

interface IBuyProduct {
  product: { id: number; quantity: number }[];
  value: { type: number; count: number }[];
  total: number;
  pay: number;
}
export const ListProductService = () => {
  return axios({
    url: ListItemsPath,
    method: 'GET'
  }) as Promise<AxiosResponse<IApiReturn<IListItems[]>, any>>;
};

export const BuyProductService = (data: IBuyProduct) => {
  return axios({
    url: BuyItems,
    method: 'POST',
    data
  }) as Promise<AxiosResponse<IApiReturn<any>, any>>;
};

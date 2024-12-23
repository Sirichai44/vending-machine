export interface IApiReturn<T> {
  status: number;
  message: string;
  results: T;
}

export interface IListItems {
  id: number;
  name: string;
  price: number;
  image_url: string;
  stock: number;
}

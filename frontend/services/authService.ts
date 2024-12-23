import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import * as EP from '@/store/services/endpoint';
import { IAuthRegister } from '../typings/auth/types';
// import { useAppSelector } from '../store';

// const profile = useAppSelector((state) => state.auth.profile);
const isProd = process.env.NODE_ENV === 'production';
const authApi: any = createApi({
  baseQuery: fetchBaseQuery({ baseUrl: isProd ? EP.prod : undefined }),
  endpoints: (builder) => ({
    login: builder.mutation({
      query: (body: { email: string; password: string }) => ({
        url: EP.Login,
        method: 'POST',
        // headers: {
        //   Authorization: `Bearer ${useAppSelector((state) => state.auth.profile.token)}`
        // },
        body
      })
    }),
    register: builder.mutation({
      query: (body: IAuthRegister) => ({
        url: EP.Register,
        method: 'POST',
        body
      })
    }),
    getUser: builder.query({
      query: () => '/auth/user'
    })
  })
});

export const { useLoginMutation, useRegisterMutation, useGetUserQuery } = authApi;
export default authApi;

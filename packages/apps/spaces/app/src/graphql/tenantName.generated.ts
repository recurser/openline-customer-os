// @ts-nocheck remove this when typscript-react-query plugin is fixed
import * as Types from '../types/__generated__/graphql.types';

import { GraphQLClient } from 'graphql-request';
import { RequestInit } from 'graphql-request/dist/types.dom';
import {
  useQuery,
  useInfiniteQuery,
  UseQueryOptions,
  UseInfiniteQueryOptions,
} from '@tanstack/react-query';

function fetcher<TData, TVariables extends { [key: string]: any }>(
  client: GraphQLClient,
  query: string,
  variables?: TVariables,
  requestHeaders?: RequestInit['headers'],
) {
  return async (): Promise<TData> =>
    client.request({
      document: query,
      variables,
      requestHeaders,
    });
}
export type TenantNameQueryVariables = Types.Exact<{ [key: string]: never }>;

export type TenantNameQuery = { __typename?: 'Query'; tenant: string };

export const TenantNameDocument = `
    query TenantName {
  tenant
}
    `;
export const useTenantNameQuery = <TData = TenantNameQuery, TError = unknown>(
  client: GraphQLClient,
  variables?: TenantNameQueryVariables,
  options?: UseQueryOptions<TenantNameQuery, TError, TData>,
  headers?: RequestInit['headers'],
) =>
  useQuery<TenantNameQuery, TError, TData>(
    variables === undefined ? ['TenantName'] : ['TenantName', variables],
    fetcher<TenantNameQuery, TenantNameQueryVariables>(
      client,
      TenantNameDocument,
      variables,
      headers,
    ),
    options,
  );
useTenantNameQuery.document = TenantNameDocument;

useTenantNameQuery.getKey = (variables?: TenantNameQueryVariables) =>
  variables === undefined ? ['TenantName'] : ['TenantName', variables];
export const useInfiniteTenantNameQuery = <
  TData = TenantNameQuery,
  TError = unknown,
>(
  pageParamKey: keyof TenantNameQueryVariables,
  client: GraphQLClient,
  variables?: TenantNameQueryVariables,
  options?: UseInfiniteQueryOptions<TenantNameQuery, TError, TData>,
  headers?: RequestInit['headers'],
) =>
  useInfiniteQuery<TenantNameQuery, TError, TData>(
    variables === undefined
      ? ['TenantName.infinite']
      : ['TenantName.infinite', variables],
    (metaData) =>
      fetcher<TenantNameQuery, TenantNameQueryVariables>(
        client,
        TenantNameDocument,
        { ...variables, ...(metaData.pageParam ?? {}) },
        headers,
      )(),
    options,
  );

useInfiniteTenantNameQuery.getKey = (variables?: TenantNameQueryVariables) =>
  variables === undefined
    ? ['TenantName.infinite']
    : ['TenantName.infinite', variables];
useTenantNameQuery.fetcher = (
  client: GraphQLClient,
  variables?: TenantNameQueryVariables,
  headers?: RequestInit['headers'],
) =>
  fetcher<TenantNameQuery, TenantNameQueryVariables>(
    client,
    TenantNameDocument,
    variables,
    headers,
  );

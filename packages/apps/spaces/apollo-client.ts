import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client';
import { setContext } from '@apollo/client/link/context';

const httpLink = new HttpLink({
  uri: `/customer-os-api/query`,
  fetchOptions: {
    credentials: 'include',
  },
});

const authLink = setContext((_, { headers }) => {
  return {
    headers: {
      ...headers,
      'Content-Type': 'application/json',
    },
  };
});

const client = new ApolloClient({
  cache: new InMemoryCache({
    typePolicies: {
      Contact: {
        fields: {
          timelineEvents: {
            keyArgs: false,
            merge(existing = [], incoming) {
              return [...incoming, ...existing];
            },
          },
        },
      },
      Organization: {
        fields: {
          timelineEvents: {
            keyArgs: false,
            merge(existing = [], incoming) {
              return [...incoming, ...existing];
            },
          },
        },
      },
      Query: {
        fields: {
          dashboardView: {
            keyArgs: false,
            merge(
              existing = { content: [] },
              incoming,
              {
                args: {
                  // @ts-expect-error look into it later
                  pagination: { page, limit },
                },
              },
            ) {
              if (page === 1) return incoming;
              return {
                ...existing,
                content: [...existing.content, ...incoming.content],
              };
            },
          },
        },
      },
    },
  }),
  link: authLink.concat(httpLink),
  queryDeduplication: true,
  assumeImmutableResults: true,
  connectToDevTools: true,
  credentials: 'include',
});

export default client;

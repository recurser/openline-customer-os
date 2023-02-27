import '@openline-ai/openline-web-chat/dist/esm/index.css';
import { Configuration, FrontendApi, Session } from '@ory/client';
import { edgeConfig } from '@ory/integrations/next';
import React, { useEffect, useState } from 'react';
import { WebChat } from '@openline-ai/openline-web-chat';
import { useRouter } from 'next/router';
import { getUserName } from '../../../../utils';
import { SidePanel } from '../../organisms';
import { PageContentLayout } from '../page-content-layout';
import client from '../../../../apollo-client';
import { ApolloProvider } from '@apollo/client';

const ory = new FrontendApi(new Configuration(edgeConfig));

export const MainPageWrapper = ({ children }: any) => {
  const router = useRouter();
  // const setTheme = (theme) => {
  //     document.documentElement.className = theme;
  //     localStorage.setItem('theme', theme);
  // }
  // const getTheme = () => {
  //     const theme = localStorage.getItem('theme');
  //     theme && setTheme(theme);
  // }

  const [session, setSession] = useState<Session | undefined>();
  const [userEmail, setUserEmail] = useState<string | undefined>();
  const [logoutUrl, setLogoutUrl] = useState<string | undefined>();

  useEffect(() => {
    if (router.asPath.startsWith('/login')) {
      return;
    }
    ory
      .toSession()
      .then(({ data }) => {
        // User has a session!
        setSession(data);
        setUserEmail(getUserName(data.identity));
        // Create a logout url
        ory.createBrowserLogoutFlow().then(({ data }) => {
          setLogoutUrl(data.logout_url);
        });
      })
      .catch(() => {
        // Redirect to login page
        return router.push(edgeConfig.basePath + '/ui/login');
      });
  }, [router]);

  if (!session) {
    if (router.asPath.startsWith('/login')) {
      return <>{children}</>;
    }
    if (router.asPath !== '/login') {
      return null;
    }
  }

  return <ApolloProvider client={client}>{children}</ApolloProvider>;
};
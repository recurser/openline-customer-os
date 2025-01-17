'use client';
import React, { ChangeEvent, useEffect, useState } from 'react';
import axios from 'axios';
import { GetServerSidePropsContext } from 'next';
import { getServerSession } from 'next-auth/next';
import { useQueryClient } from '@tanstack/react-query';
import { useRouter, useSearchParams } from 'next/navigation';

import { Card, CardBody, CardHeader } from '@ui/layout/Card';
import { signIn, useSession } from 'next-auth/react';
import { Divider } from '@ui/presentation/Divider';
import { Text } from '@ui/typography/Text';
import { Heading } from '@ui/typography/Heading';
import { Switch } from '@ui/form/Switch';
import { Flex } from '@ui/layout/Flex';
import { HStack, VStack } from '@ui/layout/Stack';
import { Spinner } from '@ui/feedback/Spinner';
import { useGlobalCacheQuery } from '@shared/graphql/global_Cache.generated';

import { FormLabel } from '@ui/form/FormElement';
import { Icons } from '@ui/media/Icon';
import {
  GetGoogleSettings,
  GetSlackSettings,
  OAuthUserSettingsInterface,
  SlackSettingsInterface,
} from 'services/settings/settingsService';
import { authOptions } from 'pages/api/auth/[...nextauth]';
import { RevokeAccess } from 'services/admin/userAdminService';
import { toastError, toastSuccess } from '@ui/presentation/Toast';

export const AuthPanel = () => {
  const router = useRouter();
  const { data: session } = useSession();
  const queryClient = useQueryClient();

  const queryParams = useSearchParams();

  useEffect(() => {
    if (
      queryParams &&
      queryParams.has('redirect_slack') &&
      queryParams.has('code')
    ) {
      setSlackSettingsLoading(true);

      axios
        .post(`/ua/slack/oauth/callback?code=${queryParams.get('code')}`)
        .then(({ data }) => {
          GetSlackSettings().then((res: SlackSettingsInterface) => {
            setSlackSettings(res);
            setSlackSettingsLoading(false);
          });
          router.push('/settings?tab=auth');
        })
        .catch((reason) => {
          router.push('/settings?tab=auth');
        });
    } else {
      setSlackSettingsLoading(true);
      GetSlackSettings().then((res: SlackSettingsInterface) => {
        setSlackSettings(res);
        setSlackSettingsLoading(false);
      });
    }
  }, [queryParams]);

  const [googleSettingsLoading, setGoogleSettingsLoading] = useState(true);
  const [googleSettings, setGoogleSettings] =
    useState<OAuthUserSettingsInterface>({
      gmailSyncEnabled: false,
      googleCalendarSyncEnabled: false,
    });

  const [slackSettingsLoading, setSlackSettingsLoading] = useState(true);
  const [slackSettings, setSlackSettings] = useState<SlackSettingsInterface>({
    slackEnabled: false,
  });

  useEffect(() => {
    if (session) {
      setGoogleSettingsLoading(true);
      GetGoogleSettings(session.user.playerIdentityId).then(
        (res: OAuthUserSettingsInterface) => {
          setGoogleSettings(res);
          setGoogleSettingsLoading(false);
        },
      );
    }
  }, [session]);

  const handleSyncGoogleToggle = async (event: ChangeEvent) => {
    setGoogleSettingsLoading(true);
    const scopes = [
      'openid',
      'email',
      'profile',
      'https://www.googleapis.com/auth/gmail.readonly',
      'https://www.googleapis.com/auth/gmail.send',
      'https://www.googleapis.com/auth/calendar.readonly',
    ];

    if ((event.target as HTMLInputElement).checked) {
      const _ = await signIn(
        'google',
        { callbackUrl: '/settings?tab=oauth' },
        {
          prompt: 'login',
          scope: scopes.join(' '),
        },
      );
    } else {
      RevokeAccess('google', {
        // @ts-expect-error look into it
        providerAccountId: session.user.playerIdentityId,
      })
        .then((data) => {
          // @ts-expect-error look into it
          GetGoogleSettings(session.user.playerIdentityId)
            .then((res: OAuthUserSettingsInterface) => {
              setGoogleSettings(res);
              setGoogleSettingsLoading(false);
              queryClient.invalidateQueries(useGlobalCacheQuery.getKey());
            })
            .catch(() => {
              setGoogleSettingsLoading(false);
              toastError(
                'There was a problem on our side and we cannot load settings data at the moment, we are doing our best to solve it! ',
                'revoke-google-access',
              );
            });
          toastSuccess(
            'We have successfully revoked the access to your google account!',
            'revoke-google-access',
          );
        })
        .catch(() => {
          setGoogleSettingsLoading(false);
          toastError(
            'There was a problem on our side and we cannot load settings data at the moment, we are doing our best to solve it! ',
            'revoke-google-access',
          );
        });
    }
  };

  const handleSlackToggle = async (event: ChangeEvent) => {
    setSlackSettingsLoading(true);

    if ((event.target as HTMLInputElement).checked) {
      axios
        .get(`/ua/slack/requestAccess`)
        .then(({ data }) => {
          location.href = data.url;
        })
        .catch((reason) => {
          toastError(
            'There was a problem on our side and we cannot load settings data at the moment, we are doing our best to solve it! ',
            'request-access-slack-access',
          );
          setSlackSettingsLoading(false);
        });
    } else {
      RevokeAccess('slack')
        .then((data) => {
          GetSlackSettings()
            .then((res: SlackSettingsInterface) => {
              setSlackSettings(res);
              setSlackSettingsLoading(false);
            })
            .catch(() => {
              setSlackSettingsLoading(false);
              toastError(
                'There was a problem on our side and we cannot load settings data at the moment, we are doing our best to solve it! ',
                'revoke-slack-access',
              );
            });
          toastSuccess(
            `We can't access your Slack workspace anymore`,
            'revoke-slack-access',
          );
        })
        .catch(() => {
          setSlackSettingsLoading(false);
          toastError(
            'There was a problem on our side and we cannot load settings data at the moment, we are doing our best to solve it! ',
            'revoke-slack-access',
          );
        });
    }
  };

  return (
    <>
      <Card
        bg='#FCFCFC'
        borderRadius='2xl'
        flexDirection='column'
        boxShadow='none'
        position='relative'
        background='gray.25'
      >
        <CardHeader px={6} pb={2}>
          <Flex gap='1' align='center' mb='2'>
            <Icons.GOOGLE boxSize='6' />
            <Heading as='h1' fontSize='lg' color='gray.700'>
              Google OAuth
            </Heading>
          </Flex>
          <Divider></Divider>
        </CardHeader>

        <CardBody padding={6} pr={0} pt={0} position='unset'>
          <Text noOfLines={2} mt={2} mb={3}>
            Enable OAuth Integration to get access to your google workspace
            emails and calendar events
          </Text>
          <Flex direction={'column'} gap={2} width={'250px'}>
            <HStack>
              <VStack alignItems={'start'}>
                <Flex gap='1' align='center'>
                  <Icons.GMAIL boxSize='6' />
                  <FormLabel mb='0'>Sync Google Mail</FormLabel>
                </Flex>
                <Flex gap='1' align='center'>
                  <Icons.GOOGLE_CALENDAR boxSize='6' />
                  <FormLabel mb='0'>Sync Google Calendar</FormLabel>
                </Flex>
              </VStack>

              {googleSettingsLoading && <Spinner size='sm' color='green.500' />}
              {!googleSettingsLoading && (
                <Switch
                  isChecked={googleSettings.gmailSyncEnabled}
                  colorScheme='green'
                  onChange={(event) => handleSyncGoogleToggle(event)}
                />
              )}
            </HStack>
          </Flex>
        </CardBody>
      </Card>

      <Card
        bg='#FCFCFC'
        borderRadius='2xl'
        flexDirection='column'
        boxShadow='none'
        position='relative'
        background='gray.25'
        mt={4}
      >
        <CardHeader px={6} pb={2}>
          <Flex gap='1' align='center' mb='2'>
            <Icons.Slack boxSize='6' />
            <Heading as='h1' fontSize='lg' color='gray.700'>
              Slack
            </Heading>
          </Flex>
          <Divider></Divider>
        </CardHeader>

        <CardBody padding={6} pr={0} pt={0} position='unset'>
          <Text noOfLines={2} mt={2} mb={3}>
            Enable Slack Integration to get access to your Slack workspace
          </Text>
          <Flex direction={'column'} gap={2} width={'250px'}>
            <HStack>
              <VStack alignItems={'start'}>
                <Flex gap='1' align='center'>
                  <Icons.Slack boxSize='6' />
                  <FormLabel mb='0'>Sync Slack</FormLabel>
                </Flex>
              </VStack>

              {slackSettingsLoading && <Spinner size='sm' color='green.500' />}
              {!slackSettingsLoading && (
                <Switch
                  isChecked={slackSettings.slackEnabled}
                  colorScheme='green'
                  onChange={(event) => handleSlackToggle(event)}
                />
              )}
            </HStack>
          </Flex>
        </CardBody>
      </Card>
    </>
  );
};

export async function getServerSideProps(context: GetServerSidePropsContext) {
  const session = await getServerSession(context.req, context.res, authOptions);
  return { props: { session } };
}

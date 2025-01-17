import { Flex } from '@ui/layout/Flex';
import { Text } from '@ui/typography/Text';
import { Card, CardBody } from '@ui/presentation/Card';
import { User, Contact } from '@graphql/types';
import React, { useCallback, useMemo } from 'react';
import { Image } from '@ui/media/Image';
import noteIcon from 'public/images/event-ill-log-stub.png';
import { Box } from '@ui/layout/Box';
import { Mail01 } from '@ui/media/icons/Mail01';
import { Calendar } from '@ui/media/icons/Calendar';
import { Phone } from '@ui/media/icons/Phone';
import { MessageTextSquare01 } from '@ui/media/icons/MessageTextSquare01';
import { LogEntryWithAliases } from '@organization/src/components/Timeline/types';
import { HtmlContentRenderer } from '@ui/presentation/HtmlContentRenderer/HtmlContentRenderer';
import { useTimelineEventPreviewMethodsContext } from '@organization/src/components/Timeline/preview/context/TimelineEventPreviewContext';

interface LogEntryStubProps {
  data: LogEntryWithAliases;
}

function getAuthor(user: User | Contact) {
  if (!user) return 'Unknown';

  if (user.name) {
    return user.name;
  }

  if (user.firstName || user.lastName) {
    return `${user.firstName ?? ''} ${user.lastName ?? ''}`.trim();
  }

  return 'Unknown';
}

export const LogEntryStub = ({ data }: LogEntryStubProps) => {
    const { openModal } = useTimelineEventPreviewMethodsContext();
  const isTemporary = !data?.updatedAt;


  const fullName = getAuthor(data?.logEntryCreatedBy);
  const getLogEntryIcon = useCallback((type: string | null) => {
    switch (type) {
      case 'email':
        return <Mail01 color='gray.500' boxSize={3} />;
      case 'meeting':
        return <Calendar color='gray.500' boxSize={3} />;
      case 'voicemail':
      case 'call':
        return <Phone color='gray.500' boxSize={3} />;
      case 'text-message':
        return <MessageTextSquare01 color='gray.500' boxSize={3} />;

      default:
        return null;
    }
  }, []);

  const getInlineTags = useCallback(() => {
    if (data.tags?.[0]?.name) {
      return data.tags?.[0]?.name;
    }
    const parser = new DOMParser();
    const doc = parser.parseFromString(`<p>${data?.content}</p>`, 'text/html');
    const element = doc.querySelector('.customeros-tag');
    // Return the inner HTML of the found element
    return element?.innerHTML || null;
  }, [data.tags, data.content]);

  const logEntryIcon = useMemo(() => {
    const firstTag = getInlineTags();
    const icon = getLogEntryIcon(firstTag);

    if (!icon) return null;
    return (
      <Flex
        zIndex={1}
        position='relative'
        bg='white'
        border='1px solid'
        borderColor='gray.200'
        borderRadius='md'
        p={2}
        right='-3px'
        top='3px'
      >
        {icon}
      </Flex>
    );
  }, [getInlineTags]);

  return (
    <Card
      variant='outline'
      size='md'
      maxWidth={549}
      ml={6}
      boxShadow='xs'
      borderColor='gray.200'
      borderRadius='lg'
      opacity={isTemporary ? 0.5 : 1}
      onClick={() => !isTemporary && openModal(data.id)}
      _hover={{ boxShadow: 'md' }}
      transition='all 0.2s ease-out'
      cursor={isTemporary ? 'progress' : 'pointer'}
    >
      <CardBody px='3' py='2'>
        <Flex
          w='full'
          justify='space-between'
          position='relative'
          h='fit-content'
        >
          <Text
            w={460}
            noOfLines={4}
            color='gray.700'
            fontSize='sm'
            height='fit-content'
          >
            <Text as='span'>{fullName}</Text>
            <Text as='span' color='gray.500' mx={1}>
              wrote
            </Text>
            <HtmlContentRenderer
              position='relative'
              zIndex={1}
              pointerEvents='none'
              showAsInlineText
              fontSize='sm'
              noOfLines={4}
              htmlContent={`${data?.content}`}
            />
          </Text>

          <Box h={86}>
            <Box position='absolute' top={-2} right={-3}>
              <Image src={noteIcon} alt='' height={94} width={124} />
            </Box>

            {logEntryIcon && logEntryIcon}
          </Box>
        </Flex>
      </CardBody>
    </Card>
  );
};

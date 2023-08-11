'use client';
import { MouseEvent, useState, useRef } from 'react';
import { useParams } from 'next/navigation';
import { useForm } from 'react-inverted-form';
import { useQueryClient } from '@tanstack/react-query';

import { Flex } from '@ui/layout/Flex';
import { IconButton } from '@ui/form/IconButton';
import { Avatar } from '@ui/media/Avatar';
import { Icons } from '@ui/media/Icon';
import { FormInput } from '@ui/form/Input';
import { FormAutoresizeTextarea } from '@ui/form/Textarea';
import { FormSelect } from '@ui/form/SyncSelect';
import { FormInputGroup } from '@ui/form/InputGroup';
import { Card, CardBody, CardHeader } from '@ui/presentation/Card';
import { useOutsideClick } from '@ui/utils';
import { Collapse } from '@ui/transitions/Collapse';
import { useDisclosure } from '@ui/utils';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { useUpdateContactMutation } from '@organization/graphql/updateContact.generated';
import { useUpdateContactRoleMutation } from '@organization/graphql/updateContactJobRole.generated';
import { useDeleteContactMutation } from '@organization/graphql/deleteContact.generated';
import { useAddContactEmailMutation } from '@organization/graphql/addContactEmail.generated';
import { useRemoveContactEmailMutation } from '@organization/graphql/removeContactEmail.generated';
import { useAddContactPhoneNumberMutation } from '@organization/graphql/addContactPhoneNumber.generated';
import { useUpdateContactPhoneNumberMutation } from '@organization/graphql/updateContactPhoneNumber.generated';
import { useRemoveContactPhoneNumberMutation } from '@organization/graphql/removeContactPhoneNumber.generated';

import { ContactFormDto, ContactForm } from './Contact.dto';
import { invalidateQuery, timezoneOptions } from '../util';
import { ConfirmDeleteDialog } from '@ui/presentation/Modal/ConfirmDeleteDialog';
import User from '@spaces/atoms/icons/User';
import { EmailValidationMessage } from '@organization/components/Tabs/panels/PeoplePanel/ContactCard/EmailValidationMessage';
import { Contact } from '@graphql/types';
import { Fade } from '@ui/transitions/Fade';

interface ContactCardProps {
  index: number;
  contact: Contact;
}

export const ContactCard = ({ contact, index }: ContactCardProps) => {
  const client = getGraphQLClient();
  const organizationId = useParams()?.id as string;
  const queryClient = useQueryClient();
  const cardRef = useRef<HTMLDivElement>(null);
  const [isExpanded, setIsExpanded] = useState(false);
  const { isOpen, onOpen, onClose } = useDisclosure();
  useOutsideClick({
    ref: cardRef,
    handler: () => setIsExpanded(false),
  });

  const data = ContactFormDto.toForm(contact);

  const formId = `contact-form-${data.id}`;

  const invalidate = () => invalidateQuery(queryClient, organizationId);
  const updateContact = useUpdateContactMutation(client, {
    onSuccess: invalidate,
  });
  const updateRole = useUpdateContactRoleMutation(client, {
    onSuccess: invalidate,
  });
  const deleteContact = useDeleteContactMutation(client, {
    onSuccess: invalidate,
  });
  const addEmail = useAddContactEmailMutation(client, {
    onSuccess: invalidate,
  });
  const removeEmail = useRemoveContactEmailMutation(client, {
    onSuccess: invalidate,
  });
  const addPhoneNumber = useAddContactPhoneNumberMutation(client, {
    onSuccess: invalidate,
  });
  const updatePhoneNumber = useUpdateContactPhoneNumberMutation(client, {
    onSuccess: invalidate,
  });
  const removePhoneNumber = useRemoveContactPhoneNumberMutation(client, {
    onSuccess: invalidate,
  });

  const toggle = (e: MouseEvent<HTMLDivElement>) => {
    if (['name', 'role', 'title'].includes((e.target as any)?.id)) {
      setIsExpanded(true);
      return;
    }
    setIsExpanded((prev) => !prev);
  };

  const prevEmail = data.email;
  const prevPhoneNumberId = data.phoneId;

  const { state } = useForm<ContactForm>({
    formId,
    defaultValues: data,
    stateReducer: (state, action, next) => {
      if (
        action.type === 'FIELD_BLUR' &&
        //@ts-expect-error payload action type fix
        !state.fields?.[action.payload.name]?.meta.pristine
      ) {
        switch (action.payload.name) {
          case 'name':
          case 'timezone':
          case 'note': {
            updateContact.mutate(ContactFormDto.toDto({ ...state.values }));
            break;
          }
          case 'company':
          case 'title':
          case 'role': {
            const key = (() => {
              const { name } = action.payload;
              if (name === 'role') return 'description';
              if (name === 'title') return 'jobTitle';
              return name;
            })();

            updateRole.mutate({
              contactId: state.values.id,
              input: {
                id: state.values.roleId,
                description: state.values.role,
                jobTitle: state.values.title,
                company: state.values.company,
                [key]: action.payload.value,
              },
            });
            break;
          }
          case 'email': {
            const newEmail = action.payload.value;
            if (!newEmail) {
              removeEmail.mutate({ contactId: data.id, email: prevEmail });
              break;
            }
            addEmail.mutate({ contactId: data.id, input: { email: newEmail } });
            break;
          }
          case 'phone': {
            const newPhoneNumber = action.payload.value;
            if (!newPhoneNumber) {
              removePhoneNumber.mutate({
                contactId: data.id,
                id: prevPhoneNumberId,
              });
              break;
            }
            if (!prevPhoneNumberId) {
              addPhoneNumber.mutate({
                contactId: data.id,
                input: { phoneNumber: newPhoneNumber },
              });
              break;
            }
            updatePhoneNumber.mutate({
              contactId: data.id,
              input: {
                id: prevPhoneNumberId,
                phoneNumber: newPhoneNumber,
              },
            });
            break;
          }
          default:
            break;
        }
      }
      return next;
    },
  });

  const handleDelete = (e: MouseEvent) => {
    e.stopPropagation();
    e.preventDefault();
    deleteContact.mutate({ contactId: data.id }, { onSuccess: onClose });
  };

  const toggleConfirmDelete = (e: MouseEvent) => {
    e.stopPropagation();
    e.preventDefault();
    onOpen();
  };

  return (
    <>
      <Card
        key={data.id}
        w='full'
        ref={cardRef}
        boxShadow={isExpanded ? 'md' : 'xs'}
        cursor='pointer'
        borderRadius='lg'
        border='1px solid'
        borderColor='gray.200'
        _hover={{
          boxShadow: 'md',
          '& > div > #confirm-button': {
            opacity: '1',
            pointerEvents: 'auto',
          },
        }}
        transition='all 0.2s ease-out'
      >
        <CardHeader
          as={Flex}
          p='4'
          pb={isExpanded ? 2 : 4}
          position='relative'
          onClick={toggle}
        >
          <Avatar
            name={state?.values?.name ?? data?.name}
            icon={
              <User
                color={'var(--chakra-colors-primary-700)'}
                height='1.8rem'
              />
            }
          />
          <Flex ml='4' flexDir='column' flex='1'>
            <FormInput
              h='6'
              name='name'
              formId={formId}
              placeholder='Name'
              color='gray.700'
              fontWeight='semibold'
            />
            <FormInput
              h='6'
              name='title'
              color='gray.500'
              formId={formId}
              placeholder='Title'
            />
            <FormInput
              h='6'
              name='role'
              color='gray.500'
              formId={formId}
              placeholder='Role'
            />
          </Flex>
          {isExpanded && (
            <IconButton
              size='xs'
              top='2'
              right='2'
              variant='ghost'
              colorScheme='gray'
              id='collapse-button'
              position='absolute'
              aria-label='Close'
              onClick={onClose}
              icon={<Icons.Check color='gray.400' boxSize='5' />}
            />
          )}

          {!isExpanded && (
            <IconButton
              size='xs'
              top='2'
              right='2'
              variant='ghost'
              color='gray.400'
              colorScheme='gray'
              _hover={{
                background: 'red.100',
                color: 'red.400',
              }}
              opacity={0}
              pointerEvents='none'
              id='confirm-button'
              position='absolute'
              aria-label='Delete contact'
              isLoading={deleteContact.isLoading}
              onClick={toggleConfirmDelete}
              icon={<Icons.Trash1 boxSize='5' />}
            />
          )}
        </CardHeader>

        <Collapse
          in={isExpanded}
          style={{ overflow: 'unset' }}
          delay={{
            exit: 2,
          }}
        >
          <Fade
            in={isExpanded}
            delay={{
              enter: 0.2,
            }}
          >
            <CardBody pt={0}>
              <FormInputGroup
                formId={formId}
                name='company'
                placeholder='Company name'
                leftElement={<Icons.Building7 color='gray.500' />}
              />
              <FormInputGroup
                formId={formId}
                name='email'
                placeholder='Email'
                leftElement={<Icons.Mail1 color='gray.500' />}
                rightElement={
                  <EmailValidationMessage
                    email={data.email}
                    validationDetails={
                      contact?.emails?.[0]?.emailValidationDetails
                    }
                  />
                }
              />
              <FormInputGroup
                formId={formId}
                name='phone'
                placeholder='Phone number'
                leftElement={<Icons.Phone2 color='gray.500' />}
              />
              <FormSelect
                formId={formId}
                isClearable
                name='timezone'
                placeholder='Timezone'
                options={timezoneOptions}
                leftElement={<Icons.Clock color='gray.500' mr='3' />}
              />
              <FormAutoresizeTextarea
                pl='30px'
                formId={formId}
                name='note'
                placeholder='Notes'
                leftElement={<Icons.File2 color='gray.500' />}
              />
            </CardBody>
          </Fade>
        </Collapse>
      </Card>
      <ConfirmDeleteDialog
        label='Delete this contact?'
        confirmButtonLabel='Delete contact'
        isOpen={isOpen}
        onClose={onClose}
        onConfirm={handleDelete}
        isLoading={deleteContact.isLoading}
      />
    </>
  );
};
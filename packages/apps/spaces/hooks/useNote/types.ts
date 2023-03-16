export {
  useGetContactCommunicationChannelsQuery,
  useGetContactPersonalDetailsQuery,
  useCreateContactMutation,
  useUpdateContactEmailMutation,
  useAddEmailToContactMutation,
  useRemoveEmailFromContactMutation,
  useUpdateContactPhoneNumberMutation,
  useRemovePhoneNumberFromContactMutation,
  useAddPhoneToContactMutation,
  useCreateOrganizationNoteMutation,
  useCreateContactNoteMutation,
  useRemoveNoteMutation,
  useUpdateNoteMutation,
  DataSource,
  useGetContactNotesQuery,
  GetOrganizationTimelineDocument,
} from '../../graphQL/__generated__/generated';
export type {
  Contact,
  GetContactCommunicationChannelsQuery,
  GetContactPersonalDetailsQuery,
  ContactInput,
  CreateContactMutation,
  Email,
  UpdateContactEmailMutation,
  NoteInput,
  CreateOrganizationNoteMutation,
  CreateContactNoteMutation,
  GetContactNotesQuery,
  RemoveNoteMutation,
  NoteUpdateInput,
  UpdateNoteMutation,
  GetOrganizationTimelineQuery,
  Note,
} from '../../graphQL/__generated__/generated';

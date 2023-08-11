import { Contact } from '@graphql/types';
import { SelectOption } from '@shared/types/SelectOptions';

import { UpdateContactMutationVariables } from '@organization/graphql/updateContact.generated';

export interface ContactForm {
  id: string;
  name: string;
  title: string;
  role: string;
  roleId: string;
  note: string;
  email: string;
  phone: string;
  phoneId: string;
  timezone: SelectOption<string> | null;
  company?: string;
}

export class ContactFormDto implements ContactForm {
  id: string; // auxiliary field
  name: string;
  role: string;
  roleId: string; // auxiliary field
  title: string;
  note: string;
  email: string;
  phone: string;
  phoneId: string; // auxiliary field
  timezone: SelectOption<string> | null;
  company: string;

  constructor(data?: Partial<Contact> | null) {
    this.id = data?.id || ''; // auxiliary field
    this.name = data?.firstName || '';
    this.title = data?.jobRoles?.[0]?.jobTitle || '';
    this.roleId = data?.jobRoles?.[0]?.id || ''; // auxiliary field
    this.role = data?.jobRoles?.[0]?.description || '';
    this.note = data?.description || '';
    this.email = data?.emails?.[0]?.email || '';
    this.phone = data?.phoneNumbers?.[0]?.rawPhoneNumber || '';
    this.phoneId = data?.phoneNumbers?.[0]?.id || ''; // auxiliary field
    this.timezone = data?.timezone
      ? { label: data?.timezone, value: data?.timezone }
      : null;
    this.company = data?.jobRoles?.[0]?.company || '';
  }

  static toForm(data: Contact) {
    return new ContactFormDto(data);
  }

  static toDto(payload: ContactForm): UpdateContactMutationVariables {
    return {
      input: {
        id: payload.id,
        firstName: payload.name,
        description: payload.note,
        prefix: payload.title,
        timezone: payload.timezone?.value,
      },
    };
  }
}
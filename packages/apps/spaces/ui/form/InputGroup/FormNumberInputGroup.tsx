'use client';
import { FC } from 'react';
import { useField } from 'react-inverted-form';
import { InputGroup, InputLeftElement, InputRightElement } from './InputGroup';
import {
  NumberInput,
  NumberInputField,
} from '@ui/form/NumberInput/NumberInput';
import {
  FormControl,
  FormLabel,
  NumberInputProps,
  VisuallyHidden,
} from '@chakra-ui/react';

interface FormNumberInputGroupProps extends NumberInputProps {
  name: string;
  formId: string;
  leftElement?: React.ReactNode;
  rightElement?: React.ReactNode;
  label?: string;
  isLabelVisible?: boolean;
}

export const FormNumberInputGroup: FC<FormNumberInputGroupProps> = ({
  name,
  formId,
  leftElement,
  rightElement,
  label,
  isLabelVisible,
  ...rest
}) => {
  const { getInputProps } = useField(name, formId);

  return (
    <FormControl>
      {isLabelVisible ? (
        <FormLabel fontWeight={600} color={rest?.color} fontSize='sm' mb={-1}>
          {label}
        </FormLabel>
      ) : (
        <VisuallyHidden>
          <FormLabel>{label}</FormLabel>
        </VisuallyHidden>
      )}
      <InputGroup>
        {leftElement && (
          <InputLeftElement w='4'>{leftElement}</InputLeftElement>
        )}

        <NumberInput {...rest} {...getInputProps()}>
          <NumberInputField
            pl='30px'
            pr={0}
            autoComplete='off'
            placeholder={rest?.placeholder || ''}
          />
        </NumberInput>
        {rightElement && <InputRightElement>{rightElement}</InputRightElement>}
      </InputGroup>
    </FormControl>
  );
};

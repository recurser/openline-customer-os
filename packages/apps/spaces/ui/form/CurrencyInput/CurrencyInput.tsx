import React, { ForwardedRef } from 'react';
import {
  FormControl,
  FormLabel,
  NumberInputProps,
  VisuallyHidden,
} from '@chakra-ui/react';
import {
  InputGroup,
  InputLeftElement,
  InputRightElement,
} from '@ui/form/InputGroup/InputGroup';
import {
  NumberInput,
  NumberInputField,
} from '@ui/form/NumberInput/NumberInput';

export interface CurrencyInputProps
  extends Omit<NumberInputProps, 'onChange' | 'value'> {
  value: string;
  onChange?: (value: string) => void;
  leftElement?: React.ReactNode;
  rightElement?: React.ReactNode;
  label?: string;
  isLabelVisible?: boolean;
  formatValue?: (val: string) => string;
  parseValue?: (val: string) => string;
}

export const CurrencyInput = React.forwardRef(
  (
    {
      isLabelVisible,
      label,
      leftElement,
      rightElement,
      value,
      onChange,
      formatValue,
      parseValue,
      ...rest
    }: CurrencyInputProps,
    ref: ForwardedRef<any>,
  ) => {
    const handleValueChange = (valueString: string) => {
      // handle weird case of blurring the field with an empty value
      if (valueString === '-9007199254740991') {
        onChange?.('');
        return;
      }
      if (parseValue) {
        onChange?.(parseValue(valueString));
        return;
      }
      onChange?.(valueString);
    };

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

          <NumberInput
            {...rest}
            value={formatValue ? formatValue(value) : value}
            onChange={handleValueChange}
            _placeholder={{ color: 'gray.600' }}
          >
            <NumberInputField
              ref={ref}
              pl={leftElement ? '30px' : '0'}
              pr={0}
              autoComplete='off'
              placeholder={rest?.placeholder || ''}
            />
          </NumberInput>

          {rightElement && (
            <InputRightElement>{rightElement}</InputRightElement>
          )}
        </InputGroup>
      </FormControl>
    );
  },
);

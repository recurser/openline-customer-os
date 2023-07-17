import { createMultiStyleConfigHelpers } from '@chakra-ui/styled-system';

const helpers = createMultiStyleConfigHelpers(['field', 'addon']);

export const Input = helpers.defineMultiStyleConfig({
  baseStyle: {
    field: {
      _placeholder: {
        color: 'gray.400',
      },
    },
  },
  variants: {
    flushed: {
      field: {
        _focus: {
          borderColor: 'teal.500',
          boxShadow: 'unset',
        },
      },
    },
  },
  defaultProps: {
    variant: 'flushed',
  },
});
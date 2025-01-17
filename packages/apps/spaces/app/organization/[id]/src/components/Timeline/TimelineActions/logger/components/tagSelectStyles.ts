import omit from 'lodash/omit';
import { StylesConfig, GroupBase, OptionProps } from 'chakra-react-select';
import { CSSWithMultiValues } from '@chakra-ui/react';
import { suggestedTags } from './TagSelect';

type ChakraStylesConfig<
  OptionType = unknown,
  GroupType extends GroupBase<OptionType> = GroupBase<OptionType>,
> = Partial<StylesConfig<OptionType, true, GroupType>> & {
  [key: string]: ((...args: any[]) => CSSWithMultiValues) | undefined;
};

export const tagsSelectStyles = (
  chakraStyles: ChakraStylesConfig<unknown, GroupBase<unknown>> | undefined,
) => ({
  multiValue: (base: CSSWithMultiValues) => ({
    ...base,
    padding: 0,
    gap: 0,
    color: 'gray.700',
    background: 'transparent',
    border: '1px solid',
    borderColor: 'transparent',
    fontSize: 'var(--tag-select-font-size)',
    margin: 0,
    marginRight: 1,
    cursor: 'text',
    fontWeight: 500,

    '&:first-of-type': {
      transform: 'translateX(100px)',
    },

    '&:before': {
      content: '"#"',
    },
  }),
  clearIndicator: (base: CSSWithMultiValues) => ({
    ...base,
    background: 'transparent',
    color: 'transparent',
    display: 'none',
  }),
  multiValueRemove: (styles: CSSWithMultiValues) => ({
    ...styles,
    display: 'none',
  }),
  container: (props: CSSWithMultiValues) => ({
    ...props,
    minWidth: '300px',
    width: '100%',
    overflow: 'visible',
    cursor: 'text',
    fontSize: 'var(--tag-select-font-size)',
    padding: 0,
    _focusVisible: { border: 'none !important' },
    _focus: { border: 'none !important' },
  }),
  menuList: (
    props: CSSWithMultiValues,
    data: { options: Array<OptionProps & { value: string }> },
  ) => {
    const isNew =
      data?.options?.length === 1 &&
      data?.options?.[0]?.label === data.options?.[0]?.value &&
      !suggestedTags.includes(data.options?.[0]?.label);

    if (isNew) {
      return {
        position: 'absolute',
        bottom: '-999999999px',
      };
    }

    return {
      ...props,
      padding: '2',
      boxShadow: 'md',
      borderColor: 'gray.200',
      borderRadius: 'lg',
      maxHeight: '150px',
      fontSize: 'inherit',
    };
  },
  option: (
    props: CSSWithMultiValues,
    { isSelected, isFocused }: { isSelected: boolean; isFocused: boolean },
  ) => ({
    ...props,
    my: '2px',
    borderRadius: 'md',
    color: 'gray.700',
    bg: isSelected ? 'primary.50' : 'white',
    boxShadow: isFocused ? 'menuOptionsFocus' : 'none',
    _hover: { bg: isSelected ? 'primary.50' : 'gray.100' },
  }),
  groupHeading: (props: CSSWithMultiValues) => ({
    ...props,
    color: 'gray.400',
    textTransform: 'uppercase',
    fontWeight: 'regular',
    fontSize: 'inherit',
  }),
  input: (props: CSSWithMultiValues) => ({
    ...props,
    color: 'gray.500',
    fontWeight: 'regular',
    cursor: 'text',
    padding: 0,
    fontSize: 'var(--tag-select-font-size)',
  }),
  valueContainer: (props: CSSWithMultiValues) => ({
    ...props,
    padding: 0,
    maxH: '86px',
    overflowY: 'auto',
    fontSize: 'inherit',
  }),
  ...omit<ChakraStylesConfig<unknown, GroupBase<unknown>>>(chakraStyles, [
    'container',
    'multiValueRemove',
    'multiValue',
    'clearIndicator',
    'menuList',
    'option',
    'groupHeading',
    'input',
    'valueContainer',
  ]),
});

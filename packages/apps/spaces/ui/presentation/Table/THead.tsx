import type { HeaderContext } from '@tanstack/react-table';

import { Flex } from '@ui/layout/Flex';
import { Icons } from '@ui/media/Icon';
import { Text } from '@ui/typography/Text';
import { IconButton } from '@ui/form/IconButton';

interface THeadProps<T extends object> extends HeaderContext<T, unknown> {
  title: string;
  subTitle?: string;
  icon?: React.ReactNode;
}

export const THead = <T extends object>({
  icon,
  title,
  header,
  subTitle,
}: THeadProps<T>) => {
  const canSort = header.column.getCanSort();
  const isSorted = header.column.getIsSorted();
  const onToggleSort = header.column.getToggleSortingHandler();

  return (
    <Flex
      w='full'
      flexDir='column'
      justify='flex-start'
      alignItems='flex-start'
    >
      <Flex align='center'>
        {icon}
        <Text fontSize='sm' fontWeight='medium' color='gray.700'>
          {title}
        </Text>
        {canSort && (
          <IconButton
            ml='1'
            size='xs'
            variant='ghost'
            aria-label='Sort'
            onClick={onToggleSort}
            icon={
              !isSorted ? (
                <Icons.ArrowsSwitchVertical1 color='gray.400' />
              ) : isSorted === 'asc' ? (
                <Icons.ArrowUp color='gray.600' />
              ) : (
                <Icons.ArrowDown color='gray.600' />
              )
            }
          />
        )}
      </Flex>
      {subTitle && (
        <Text fontSize='xs' color='gray.500'>
          {subTitle}
        </Text>
      )}
    </Flex>
  );
};

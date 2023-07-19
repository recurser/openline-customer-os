'use client';

import { Flex } from '@ui/layout/Flex';

export const TabsContainer = ({ children }: { children?: React.ReactNode }) => {
  return (
    <Flex
      w='full'
      h='100%'
      bg='white'
      flexDir='column'
      border='1px solid'
      borderColor='gray.200'
      borderRadius='2xl'
    >
      {children}
    </Flex>
  );
};

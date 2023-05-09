import React, { useRef } from 'react';
import { IconButton, Image } from '../../../atoms';

export const UploadImageButton = ({ onFileChange }: { onFileChange: any }) => {
  const inputRef = useRef<HTMLInputElement | null>(null);

  const handleUploadClick = () => {
    inputRef.current?.click();
  };

  return (
    <>
      <IconButton
        id='custom-button'
        type={'button'}
        aria-label='Insert picture'
        onClick={handleUploadClick}
        isSquare
        mode='text'
        size='xxxxs'
        style={{ padding: '2px', background: 'transparent' }}
        icon={
          <Image
            color='#ccc'
            style={{ transform: 'scale(0.8)', height: '24px' }}
          />
        }
      ></IconButton>
      <input
        type='file'
        ref={inputRef}
        onChange={onFileChange}
        style={{ display: 'none' }}
      />
    </>
  );
};
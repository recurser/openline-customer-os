import React from 'react';
import Image, { StaticImageData } from 'next/image';
import styles from './avatar.module.scss';
import { getInitialsColor } from './utils';
import classNames from 'classnames';

interface AvatarProps {
  name: string;
  surname: string;
  size: number;
  image?: StaticImageData;
  imageHeight?: number;
  imageWidth?: number;
  isSquare?: boolean;
}

export const Avatar: React.FC<AvatarProps> = ({
  name,
  surname,
  size,
  image,
  imageWidth,
  imageHeight,
  isSquare = false,
  ...rest
}) => {
  if (image) {
    return (
      <>
        <Image
          {...rest}
          src={image}
          alt={`${name} ${surname}`}
          height={imageHeight || 40}
          width={imageWidth}
        />
      </>
    );
  }

  const initials = `${name?.charAt(0)}${surname?.charAt(0)}`;
  const color = getInitialsColor(initials);

  const avatarStyle = {
    width: `${size}px`,
    height: `${size}px`,
    backgroundColor: color,
    fontSize: size > 40 ? 'var(--font-size-lg)' : 'ar(--font-size-xxs)',
  };

  return (
    <div
      className={classNames(styles.avatar, {
        [styles.square]: isSquare,
      })}
      style={avatarStyle}
    >
      {initials}
    </div>
  );
};

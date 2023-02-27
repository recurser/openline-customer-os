import * as React from 'react';
import { SVGProps } from 'react';
const SvgDiscord = (props: SVGProps<SVGSVGElement>) => (
  <svg
    width={24}
    height={24}
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M14.29 11.6a.82.82 0 1 1-.82-.89.848.848 0 0 1 .82.89Zm-3.74-.89a.89.89 0 1 0 .82.89.85.85 0 0 0-.82-.89ZM19 5.65V20c-2-1.78-1.37-1.19-3.71-3.37l.42 1.48H6.64A1.639 1.639 0 0 1 5 16.46V5.65A1.65 1.65 0 0 1 6.64 4h10.72A1.65 1.65 0 0 1 19 5.65Zm-2.28 7.58a10.73 10.73 0 0 0-1.15-4.66 3.9 3.9 0 0 0-2.25-.84l-.11.13a5.25 5.25 0 0 1 2 1 6.81 6.81 0 0 0-6-.23l-.47.23a5.51 5.51 0 0 1 2.11-1l-.08-.09a3.9 3.9 0 0 0-2.25.84 10.73 10.73 0 0 0-1.15 4.66 2.9 2.9 0 0 0 2.44 1.22l.53-.67A2.5 2.5 0 0 1 9 12.84l.33.2a5.83 5.83 0 0 0 5 .28c.323-.123.631-.28.92-.47a2.55 2.55 0 0 1-1.45 1l.53.65a2.93 2.93 0 0 0 2.45-1.22l-.06-.05Z'
      fill='currentColor'
    />
  </svg>
);
export default SvgDiscord;
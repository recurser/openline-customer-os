import * as React from 'react';
import { SVGProps } from 'react';
const SvgSearch = (props: SVGProps<SVGSVGElement>) => (
  <svg
    width={24}
    height={24}
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g fill='currentColor'>
      <path d='M10.77 18.3a7.53 7.53 0 1 1 0-15.06 7.53 7.53 0 0 1 0 15.06Zm0-13.55a6 6 0 1 0 0 12 6 6 0 0 0 0-12Z' />
      <path d='M20 20.75a.74.74 0 0 1-.53-.22l-4.13-4.13a.75.75 0 0 1 1.06-1.06l4.13 4.13a.75.75 0 0 1-.53 1.28Z' />
    </g>
  </svg>
);
export default SvgSearch;
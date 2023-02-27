import * as React from 'react';
import { SVGProps } from 'react';
const SvgVolumeDown = (props: SVGProps<SVGSVGElement>) => (
  <svg
    width={24}
    height={24}
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g fill='currentColor'>
      <path d='M15 19.75a.81.81 0 0 1-.47-.16l-4.79-3.84H5a.76.76 0 0 1-.75-.75V9A.76.76 0 0 1 5 8.25h4.74l4.79-3.84a.75.75 0 0 1 1.22.59v14a.76.76 0 0 1-.43.68.71.71 0 0 1-.32.07Zm-9.25-5.5H10a.78.78 0 0 1 .47.16l3.78 3V6.56l-3.78 3a.78.78 0 0 1-.47.16H5.75v4.53ZM18.11 15.38a.75.75 0 0 1-.6-1.2 3.6 3.6 0 0 0 0-4.36.751.751 0 0 1 1.2-.9 5.07 5.07 0 0 1 0 6.16.77.77 0 0 1-.6.3Z' />
    </g>
  </svg>
);
export default SvgVolumeDown;
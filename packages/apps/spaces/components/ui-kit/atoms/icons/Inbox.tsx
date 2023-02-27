import * as React from 'react';
import { SVGProps } from 'react';
const SvgInbox = (props: SVGProps<SVGSVGElement>) => (
  <svg
    width={24}
    height={24}
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g fill='currentColor'>
      <path d='M18 20H6a2.75 2.75 0 0 1-2.75-2.75v-6a.75.75 0 1 1 1.5 0v6A1.25 1.25 0 0 0 6 18.53h12a1.25 1.25 0 0 0 1.25-1.25v-6a.75.75 0 1 1 1.5 0v6A2.75 2.75 0 0 1 18 20Z' />
      <path d='M12 15.25A3.74 3.74 0 0 1 8.29 12H4a.75.75 0 0 1-.64-1.15l3.73-6A1.69 1.69 0 0 1 8.62 4h6.76a1.74 1.74 0 0 1 1.57 1l3.69 5.94A.75.75 0 0 1 20 12h-4.29A3.74 3.74 0 0 1 12 15.25Zm-6.65-4.72H9a.75.75 0 0 1 .75.75v.22a2.25 2.25 0 0 0 4.5 0v-.22a.75.75 0 0 1 .75-.75h3.65l-3-4.86c-.08-.14-.16-.2-.26-.2H8.62a.27.27 0 0 0-.23.14l-3.04 4.92Z' />
    </g>
  </svg>
);
export default SvgInbox;
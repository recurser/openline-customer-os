import * as React from 'react';
import { SVGProps } from 'react';
const SvgFolderOpen = (props: SVGProps<SVGSVGElement>) => (
  <svg
    width={24}
    height={24}
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g fill='currentColor'>
      <path d='M4.25 18.5h-1.5v-11a2.71 2.71 0 0 1 2.68-2.75h2.41a.76.76 0 0 1 .58.25l2.67 3.23H16A2.71 2.71 0 0 1 18.72 11v.5h-1.5V11A1.219 1.219 0 0 0 16 9.75h-5.27a.74.74 0 0 1-.57-.27L7.49 6.25H5.43A1.22 1.22 0 0 0 4.25 7.5v11Z' />
      <path d='M17.12 19.25H3.5a.76.76 0 0 1-.64-.36.75.75 0 0 1 0-.74l3.38-6.5a.77.77 0 0 1 .67-.4H20.5a.76.76 0 0 1 .64.36.75.75 0 0 1 0 .74l-3.38 6.5a.77.77 0 0 1-.64.4Zm-12.39-1.5h11.94l2.6-5H7.33l-2.6 5Z' />
    </g>
  </svg>
);
export default SvgFolderOpen;
import * as React from 'react';
import { SVGProps } from 'react';
const SvgPaperclip = (props: SVGProps<SVGSVGElement>) => (
  <svg
    width={24}
    height={24}
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M8.94 20.74a5.85 5.85 0 0 1-4-1.56 5.23 5.23 0 0 1 0-7.68l7.56-7.14a4.22 4.22 0 0 1 5.69 0A4.1 4.1 0 0 1 19.5 7.3a3.46 3.46 0 0 1-1.1 2.55L10.83 17a2.47 2.47 0 0 1-3.36 0 2.23 2.23 0 0 1 0-3.28l7-6.59a.75.75 0 0 1 1.06 1.06l-7 6.59a.73.73 0 0 0 0 1.1 1 1 0 0 0 1.3 0l7.57-7.13A2.002 2.002 0 0 0 18 7.3a2.57 2.57 0 0 0-.84-1.84 2.67 2.67 0 0 0-3.63 0L6 12.59a3.729 3.729 0 0 0 0 5.5 4.4 4.4 0 0 0 6 0L19.49 11a.74.74 0 0 1 1.06 0 .75.75 0 0 1 0 1.06L13 19.18a5.84 5.84 0 0 1-4.06 1.56Z'
      fill='currentColor'
    />
  </svg>
);
export default SvgPaperclip;
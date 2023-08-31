import * as React from 'react';
import { SVGProps } from 'react';
const SvgQuote = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 16 16'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    width='1em'
    height='1em'
    {...props}
  >
    <path
      d='M4.594 7.485h.218c.594 0 1.125.22 1.594.66.48.43.719.97.719 1.62 0 .64-.24 1.17-.719 1.59-.468.41-1.052.615-1.75.615-.76 0-1.396-.29-1.906-.87-.5-.59-.75-1.425-.75-2.505 0-.99.219-1.915.656-2.775.448-.87.99-1.555 1.625-2.055C4.927 3.255 5.5 3 6 3c.333 0 .594.095.781.285.188.19.282.435.282.735 0 .44-.204.765-.61.975-.552.29-.99.63-1.312 1.02-.323.39-.506.88-.547 1.47Zm6.906 0h.203c.584 0 1.11.22 1.578.66.48.43.719.97.719 1.62 0 .63-.24 1.16-.719 1.59-.469.43-1.052.645-1.75.645-.77 0-1.406-.3-1.906-.9-.5-.6-.75-1.435-.75-2.505 0-.77.13-1.495.39-2.175.261-.68.594-1.27 1-1.77.417-.51.86-.91 1.329-1.2.479-.3.906-.45 1.281-.45.344 0 .604.095.781.285.188.18.281.425.281.735 0 .45-.203.775-.609.975-.573.31-1.005.655-1.297 1.035-.291.37-.469.855-.531 1.455Z'
      fill='currentColor'
    />
  </svg>
);
export default SvgQuote;
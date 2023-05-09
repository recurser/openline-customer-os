import * as React from 'react';
import { SVGProps } from 'react';
const SvgArchive = (props: SVGProps<SVGSVGElement>) => (
  <svg width={800} height={600} xmlns='http://www.w3.org/2000/svg' {...props}>
    <path
      stroke='#fff'
      d='M339.002 244.93c-.12 0-.225.04-.307.133a.456.456 0 0 0-.134.321v10.884c0 .118.044.22.134.298a.432.432 0 0 0 .307.118h43.996c.12 0 .225-.04.307-.118a.381.381 0 0 0 .135-.298v-10.884a.456.456 0 0 0-.135-.321.389.389 0 0 0-.307-.133h-43.996 0zm2.191 11.782a.343.343 0 0 0-.246.117.396.396 0 0 0-.105.251v29.63c0 .083.037.173.105.25.074.08.157.11.246.11h39.614c.09 0 .172-.03.246-.11a.387.387 0 0 0 .105-.25v-29.63a.39.39 0 0 0-.105-.25.342.342 0 0 0-.246-.118h-39.614 0zm13.8 7.051h12.013c.644 0 1.19.235 1.638.706a2.4 2.4 0 0 1 .68 1.723h0c0 .674-.223 1.254-.68 1.748-.448.486-.995.736-1.638.736h-12.012c-.644 0-1.19-.25-1.638-.736a2.482 2.482 0 0 1-.68-1.748h0a2.4 2.4 0 0 1 .68-1.723c.448-.471.994-.706 1.638-.706h0z'
      strokeMiterlimit={2.613}
      strokeLinejoin='round'
      strokeLinecap='round'
      strokeWidth={6.009}
      fill='none'
    />
  </svg>
);
export default SvgArchive;
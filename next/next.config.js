/** @type {import('next').NextConfig} */
const nextConfig = {
  // async headers() {
  //   console.log("headers() function called called called!!!!!!!");
  //   return [
  //     {
  //       /**
  //        * https://nextjs.org/docs/app/api-reference/next-config-js/headers#cache-control
  //        *
  //        * You cannot set Cache-Control headers in next.config.js for pages or assets,
  //        * as these headers will be overwritten in production to ensure that responses and static assets are cached effectively.
  //        *
  //        * ... however, for dev only this works!??? this is to prevent flickering on mobile due to CSS being loaded every time upon query string change
  //        */
  //       source: "/:all*(page.css)",
  //       headers: [
  //         {
  //           key: "Cache-Control",
  //           value: "public, max-age=60, immutable",
  //         },
  //       ],
  //     },
  //   ];
  // },
};

module.exports = nextConfig;

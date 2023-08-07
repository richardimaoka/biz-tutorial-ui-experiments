// https://fontawesome.com/docs/web/use-with/react/use-with#next-js
// The react-fontawesome component integrates well with Next.js
// but there is one caveat you need to solve.
// Since Next.js manages CSS differently than most web projects
// if you just follow the plain vanilla documentation to integrate
// react-fontawesome into your project you'll see huge icons
// because they are missing the accompanying CSS that makes them behave.
import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";
config.autoAddCss = false;

import "./globals.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });
import { Noto_Sans_JP } from "next/font/google";

const notojp = Noto_Sans_JP({
  weight: "400",
  preload: false,
  display: "block",
  adjustFontFallback: false,
});

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={notojp.className}>{children}</body>
    </html>
  );
}

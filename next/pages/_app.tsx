import "../styles/globals.css";
import type { AppProps } from "next/app";
import "../styles/prism-vsc-dark-plus.css";
export default function App({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />;
}

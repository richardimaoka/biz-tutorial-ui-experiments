import { ApolloProvider } from "@apollo/client";
import type { AppProps } from "next/app";
import { client } from "../libs/apolloClient";
import "../styles/globals.css";
import "../styles/prism-line-highlight.css";
import "../styles/prism-line-numbers.css";
import "../styles/prism-vsc-dark-plus.css";

import { Noto_Sans_JP } from "next/font/google";

const notojp = Noto_Sans_JP({
  weight: "400",
  preload: false,
});

export default function App({ Component, pageProps }: AppProps) {
  return (
    <ApolloProvider client={client}>
      <div className={notojp.className}>
        <Component {...pageProps} />
      </div>
    </ApolloProvider>
  );
}

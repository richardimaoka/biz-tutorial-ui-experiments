import "../styles/globals.css";
import type { AppProps } from "next/app";
import "../styles/prism-vsc-dark-plus.css";
import { ApolloProvider } from "@apollo/client";
import { client } from "../libs/apolloClient";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <ApolloProvider client={client}>
      <Component {...pageProps} />
    </ApolloProvider>
  );
}

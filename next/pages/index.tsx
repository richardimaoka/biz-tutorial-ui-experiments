import { css } from "@emotion/react";
import { GetServerSideProps } from "next";
import Link from "next/link";
import { ParsedUrlQuery } from "querystring";
import { useState } from "react";
import { Header } from "../components/Header";
import { SourceCodeViewer } from "../components/sourcecode/SourceCodeViewer";
import { TerminalComponent } from "../components/terminal/TerminalComponent";
import { client } from "../libs/apolloClient";
import { graphql } from "../libs/gql";
import { IndexSsrPageQuery } from "../libs/gql/graphql";

const queryDefinition = graphql(/* GraphQL */ `
  query IndexSsrPage($step: String, $openFilePath: String) {
    pageState(step: $step) {
      nextStep
      prevStep
      step
      sourceCode {
        ...SourceCodeViewer_Fragment
      }
      terminals {
        name
        currentDirectory
        ...TerminalComponent_Fragment
      }
    }
  }
`);

interface PageParams extends ParsedUrlQuery {
  step: string;
  openFilePath: string;
}

const extractString = (
  queryString: string | string[] | undefined
): string | undefined => {
  // if object, it must be string[]
  if (typeof queryString == "object") {
    if (queryString.length > 0) {
      return queryString[0];
    } else {
      return undefined;
    }
  } else if (typeof queryString == "string") {
    return queryString;
  } else {
    return undefined;
  }
};

export const getServerSideProps: GetServerSideProps<
  IndexSsrPageQuery,
  PageParams
> = async (context) => {
  const step = extractString(context.query.step);
  const openFilePath = extractString(context.query.openFilePath);

  const { data } = await client.query({
    query: queryDefinition,
    variables: { step, openFilePath },
  });
  return {
    props: data,
  };
};

export default function Home({ pageState }: IndexSsrPageQuery) {
  const currentPage = pageState;
  const nextStep = currentPage?.nextStep;
  const prevStep = currentPage?.prevStep;

  const terminals = currentPage?.terminals;
  const [currentTerminalIndex] = useState(0);
  const currentTerminal = terminals && terminals[currentTerminalIndex];
  const currentDirectory = currentTerminal?.currentDirectory
    ? currentTerminal.currentDirectory
    : undefined;

  return (
    currentPage && (
      <>
        <Header />
        <main
          css={css`
            background-color: #333333;
          `}
        >
          <div
            css={css`
              width: 680px;
              /* margin: 0 auto; */
              background-color: white;
            `}
          >
            {currentPage?.sourceCode && currentPage?.step && (
              <SourceCodeViewer
                fragment={currentPage.sourceCode}
                step={currentPage.step}
                currentDirectory={currentDirectory}
              />
            )}
            {currentTerminal && (
              <TerminalComponent fragment={currentTerminal} />
            )}
            {prevStep && (
              <button type="button">
                <Link href={`./?step=${prevStep}`}>prev step</Link>
              </button>
            )}
            {nextStep && (
              <button type="button">
                <Link href={`./?step=${nextStep}`}>next step</Link>
              </button>
            )}
          </div>
        </main>
      </>
    )
  );
}

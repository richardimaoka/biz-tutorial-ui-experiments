import { css } from "@emotion/react";
import { GetServerSideProps } from "next";
import { useRouter } from "next/router";
import { ParsedUrlQuery } from "querystring";
import { useState } from "react";
import { Header } from "../components/Header";
import { SourceCodeViewer } from "../components/sourcecode/SourceCodeViewer";
import { NextStepButton } from "../components/steps/NextStepButton";
import { PrevStepButton } from "../components/steps/PrevStepButton";
import { TerminalComponent } from "../components/terminal/TerminalComponent";
import { client } from "../libs/apolloClient";
import { graphql } from "../libs/gql";
import { IndexSsrPageQuery } from "../libs/gql/graphql";
import { MarkdownPane } from "../components/markdown/MarkdownPane";

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
  const router = useRouter();
  const currentStep = extractString(router.query.step);
  const openFilePath = extractString(router.query.openFilePath);

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
              display: flex;
              gap: 40px;
            `}
          >
            <div
              css={css`
                width: 680px;
              `}
            >
              {currentPage?.sourceCode && currentStep && (
                <SourceCodeViewer
                  fragment={currentPage.sourceCode}
                  step={currentStep}
                  currentDirectory={currentDirectory}
                />
              )}
              {currentTerminal && (
                <TerminalComponent fragment={currentTerminal} />
              )}
            </div>

            <div
              css={css`
                width: 680px;
              `}
            >
              <MarkdownPane />
            </div>
          </div>

          <div>
            {prevStep && (
              <PrevStepButton prevStep={prevStep} openFilePath={openFilePath} />
            )}
            {nextStep && (
              <NextStepButton nextStep={nextStep} openFilePath={openFilePath} />
            )}
          </div>
        </main>
      </>
    )
  );
}

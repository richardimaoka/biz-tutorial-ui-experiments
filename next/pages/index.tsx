import { css } from "@emotion/react";
import { GetServerSideProps } from "next";
import { useRouter } from "next/router";
import { ParsedUrlQuery } from "querystring";
import { useEffect, useState } from "react";
import { Header } from "../components/Header";
import { SourceCodeViewer } from "../components/sourcecode/SourceCodeViewer";
import { NextStepButton } from "../components/steps/NextStepButton";
import { PrevStepButton } from "../components/steps/PrevStepButton";
import { TerminalComponent } from "../components/terminal/TerminalComponent";
import { client } from "../libs/apolloClient";
import { graphql } from "../libs/gql";
import { IndexSsrPageQuery } from "../libs/gql/graphql";
import { ColumnWrapper } from "../components/column/ColumnWrapper";

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

const constructQueryString = (
  step: string | undefined,
  openFilePath: string | undefined
): string => {
  if (step) {
    if (openFilePath) {
      return `?step=${step}&openFilePath=${encodeURIComponent(openFilePath)}`;
    } else {
      return `?step=${step}`;
    }
  } else {
    if (openFilePath) {
      return `?openFilePath=${encodeURIComponent(openFilePath)}`;
    } else {
      return "";
    }
  }
};

const queryDefinition = graphql(/* GraphQL */ `
  query IndexSsrPage($tutorial: String!, $step: String) {
    page(tutorial: $tutorial, step: $step) {
      __typename
      step
      nextStep
      prevStep
      columns {
        ...ColumnWrapperFragment
      }
    }
  }
`);

interface PageParams extends ParsedUrlQuery {
  step: string;
}

export const getServerSideProps: GetServerSideProps<
  IndexSsrPageQuery,
  PageParams
> = async (context) => {
  const step = extractString(context.query.step);

  const { data } = await client.query({
    query: queryDefinition,
    variables: {
      tutorial: "sign-in-with-google",
      step: step,
    },
  });

  return {
    props: data,
  };
};

export default function Home({ page }: IndexSsrPageQuery) {
  const router = useRouter();
  const currentStep = extractString(router.query.step);
  const openFilePath = extractString(router.query.openFilePath);

  const currentPage = page;
  const nextStep = currentPage?.nextStep;
  const prevStep = currentPage?.prevStep;

  // const terminals = currentPage?.terminals;
  // const [currentTerminalIndex] = useState(0);
  // const currentTerminal = terminals && terminals[currentTerminalIndex];
  // const currentDirectory = currentTerminal?.currentDirectory
  //   ? currentTerminal.currentDirectory
  //   : undefined;

  // useEffect(() => {
  //   const handleKeyDown = (event: KeyboardEvent) => {
  //     if (event.code === "Space") {
  //       if (event.shiftKey) {
  //         prevStep &&
  //           router.push(`/${constructQueryString(prevStep, openFilePath)}`);
  //       } else {
  //         nextStep &&
  //           router.push(`/${constructQueryString(nextStep, openFilePath)}`);
  //       }
  //     }
  //   };
  //   document.addEventListener("keydown", handleKeyDown);

  //   // Don't forget to clean up
  //   return function cleanup() {
  //     document.removeEventListener("keydown", handleKeyDown);
  //   };
  // }, [router, nextStep, prevStep, openFilePath]);

  return (
    <>
      {currentPage?.columns &&
        currentPage.columns.map(
          (col, index) => col && <ColumnWrapper key={index} fragment={col} />
        )}
    </>
  );
}

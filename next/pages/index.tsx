import { css } from "@emotion/react";
import { GetServerSideProps } from "next";
import { useRouter } from "next/router";
import { ParsedUrlQuery } from "querystring";
import { useEffect, useState } from "react";
import { Header } from "../components/Header";
import { client } from "../libs/apolloClient";
import { graphql } from "../libs/gql";
import { IndexSsrPageQuery } from "../libs/gql/graphql";
import { ColumnWrapper } from "../components/column/ColumnWrapper";
import { NextButton } from "../components/navigation/NextButton";
import { PrevButton } from "../components/navigation/PrevButton";
import { AutoPlayButton } from "../components/navigation/AutoPlayButton";
import { print } from "graphql";
import { StepDisplay } from "../components/steps/StepDisplay";

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
  query IndexSsrPage($tutorial: String!, $step: String, $openFilePath: String) {
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
  const openFilePath = extractString(context.query.openFilePath);

  const { data } = await client.query({
    query: queryDefinition,
    variables: {
      tutorial: "sign-in-with-google",
      step: step,
      openFilePath: openFilePath,
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

  //console.log(print(queryDefinition));

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

  const list = [1, 2, 3, 4, 5, 6, 7, 8];

  return (
    <>
      {currentStep && <StepDisplay step={currentStep} />}
      <div
        css={css`
          display: flex;
          gap: 20px;
          height: 100vh;
          @media (max-width: 768px) {
            width: 100vw;
            height: 100vh;
          }
          width: ${768 * 2}px;
          height: 100vh;

          scroll-snap-type: x mandatory;
          scroll-behavior: smooth;
          overflow-x: auto;
        `}
      >
        {list.map((item, index) => (
          <div
            key={index}
            css={css`
              background-color: white;
              color: black;
              @media (max-width: 768px) {
                width: 100vw;
                height: 100vh;
              }
              width: 768px;
              height: 100vh;
              scroll-snap-align: start;
              flex-shrink: 0;
              padding-top: 100px;
              padding-left: 100px;
            `}
          >
            {item}
          </div>
        ))}
        {/* {page?.columns &&
          page.columns.map(
            (col, index) => col && <ColumnWrapper key={index} fragment={col} />
          )}
        {page?.columns &&
          page.columns.map(
            (col, index) => col && <ColumnWrapper key={index} fragment={col} />
          )} */}
      </div>

      {page?.prevStep && <PrevButton href={`?step=${page.prevStep}`} />}
      {page?.nextStep && <AutoPlayButton />}
      {page?.nextStep && <NextButton href={`?step=${page.nextStep}`} />}
    </>
  );
}

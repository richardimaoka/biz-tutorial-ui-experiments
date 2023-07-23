import { GetServerSideProps } from "next";
import { useRouter } from "next/router";
import { ParsedUrlQuery } from "querystring";
import { PageColumns } from "../components/column/PageColumns";
import { AutoPlayButton } from "../components/navigation/AutoPlayButton";
import { NextButton } from "../components/navigation/NextButton";
import { PrevButton } from "../components/navigation/PrevButton";
import { StepDisplay } from "../components/steps/StepDisplay";
import { client } from "../libs/apolloClient";
import { graphql } from "../libs/gql";
import { IndexSsrPageQuery } from "../libs/gql/graphql";
import { useEffect } from "react";
import { queryParamToString } from "../libs/queryString";

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
      ...PageColumnsFragment
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
  const step = queryParamToString(context.query.step);
  const openFilePath = queryParamToString(context.query.openFilePath);

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
  // const router = useRouter();

  // to debug query
  // console.log(print(queryDefinition));

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
      {page?.step && <StepDisplay step={page.step} />}
      {page && <PageColumns fragment={page} />}
      {page?.prevStep && <PrevButton href={`?step=${page.prevStep}`} />}
      {page?.nextStep && <AutoPlayButton nextStep={page.nextStep} />}
      {page?.nextStep && <NextButton href={`?step=${page.nextStep}`} />}
    </>
  );
}

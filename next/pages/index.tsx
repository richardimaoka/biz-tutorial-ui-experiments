import { GetServerSideProps } from "next";
import { useRouter } from "next/router";
import { ParsedUrlQuery } from "querystring";
import { PageColumns } from "../components/column/PageColumns";
import { AutoPlayButton } from "../components/navigation/AutoPlayButton";
import { NextButton } from "../components/navigation/NextButton";
import { PrevButton } from "../components/navigation/PrevButton";
import { StepDisplay } from "../components/navigation/StepDisplay";
import { client } from "../libs/apolloClient";
import { graphql } from "../libs/gql";
import { IndexSsrPageQuery } from "../libs/gql/graphql";
import { useEffect } from "react";
import { queryParamToString } from "../libs/queryString";
import { Header } from "../components/Header";
import { BackToStart } from "../components/navigation/BackToStart";
import { Navigation } from "../components/navigation/Navigation";

const queryDefinition = graphql(/* GraphQL */ `
  query IndexSsrPage($tutorial: String!, $step: String, $openFilePath: String) {
    page(tutorial: $tutorial, step: $step) {
      __typename
      ...NavigationFragment
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
      {page && <PageColumns fragment={page} />}
      {page && <Navigation fragment={page} />}
    </>
  );
}

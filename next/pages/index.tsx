import { css } from "@emotion/react";
import { GetServerSideProps } from "next";
import Link from "next/link";
import { useRouter } from "next/router";
import { ParsedUrlQuery } from "querystring";
import { useState } from "react";
import { Header } from "../components/Header";
import { SourceCodeViewer } from "../components/sourcecode/SourceCodeViewer";
import { TerminalComponent } from "../components/terminal/TerminalComponent";
import { client } from "../libs/apolloClient";
import { graphql } from "../libs/gql";
import { IndexSsrPageQuery } from "../libs/gql/graphql";

const queryDefinition = graphql(/* GraphQL */ `
  query IndexSsrPage($step: String) {
    pageState(step: $step) {
      nextStep
      prevStep
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
}

export const getServerSideProps: GetServerSideProps<
  IndexSsrPageQuery,
  PageParams
> = async (context) => {
  const step = typeof context.query.step == "string" ? context.query.step : "";
  const { data } = await client.query({
    query: queryDefinition,
    variables: { step },
  });
  return {
    props: data,
  };
};

export default function Home({ pageState }: IndexSsrPageQuery) {
  const router = useRouter();
  const { step } = router.query;
  const stepVariable = typeof step === "string" ? step : undefined;

  const currentPage = pageState;
  const nextStep = currentPage?.nextStep;
  const prevStep = currentPage?.prevStep;

  // CONSIDERING HOW TO IMPLEMENT SERVER-SIDE defaultOpenFile...
  //
  // 1. useState style
  //   this is not going to work, as rendering always triggered twice,
  //   by page's GraphQL query -> setState
  // x const [openFile, setOpenFile] = useState<OpenFile>(null);
  //
  // const defaultOpenFile = data?.defaultOpenFile
  //
  // 2. useMemo style
  //    this is not going to work either, as useMemo is called only at rendering
  //    so something e.g. useState should trigger rendering
  // x const openFile = useMemo(() => {
  // }, [defaultOpenFile, openFilePath]);
  //
  // 3. server-side rendering style, where step and defaultOpenFile are query params
  //    maybe this works??? declaretive style?
  //    using <Link to=="" /> avoids communication to server?

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
            {currentPage?.sourceCode && (
              <SourceCodeViewer
                fragment={currentPage.sourceCode}
                step={stepVariable}
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

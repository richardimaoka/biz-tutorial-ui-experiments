import { useQuery } from "@apollo/client";
import { css } from "@emotion/react";
import Link from "next/link";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { Header } from "../components/Header";
import { SourceCodeViewer } from "../components/sourcecode/SourceCodeViewer";
import { TerminalComponent } from "../components/terminal/TerminalComponent";
import { graphql } from "../libs/gql";
import { nonNullArray } from "../libs/nonNullArray";

const PageQuery = graphql(/* GraphQL */ `
  query PageQuery($step: String) {
    pageState(step: $step) {
      nextStep
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

export default function Home() {
  const router = useRouter();
  const { step } = router.query;
  const stepVariable = typeof step === "string" ? step : undefined;

  const { loading, error, data, client } = useQuery(PageQuery, {
    variables: { step: stepVariable },
  });

  const currentPage = data?.pageState;
  const nextStep = currentPage?.nextStep;

  const terminals = currentPage?.terminals;
  const [currentTerminalIndex] = useState(0);
  const currentTerminal = terminals && terminals[currentTerminalIndex];
  const currentDirectory = currentTerminal?.currentDirectory
    ? nonNullArray(currentTerminal?.currentDirectory)
    : undefined;

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.code === "Space" && nextStep) {
        router.push(`./?step=${nextStep}`);
      }
    };
    document.addEventListener("keyup", handleKeyDown);

    // Don't forget to clean up
    return function cleanup() {
      document.removeEventListener("keyup", handleKeyDown);
    };
  }, [step, nextStep]);

  // Page load optimization:
  useEffect(() => {
    if (nextStep) {
      client
        .query({
          query: PageQuery,
          variables: { step: nextStep },
        })
        .catch((error) => console.log(error));
    }
  }, [nextStep]);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error : {error.message}</p>;

  return (
    data && (
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
              margin: 0 auto;
              background-color: white;
            `}
          >
            {currentPage?.sourceCode && (
              <SourceCodeViewer
                fragment={currentPage.sourceCode}
                currentDirectory={currentDirectory}
              />
            )}
            {currentTerminal && (
              <TerminalComponent fragment={currentTerminal} />
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

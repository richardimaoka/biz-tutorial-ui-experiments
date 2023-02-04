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

  const [currentTerminalIndex] = useState(0);
  const terminals = data?.pageState?.terminals;

  const currentTerminal = terminals && terminals[currentTerminalIndex];
  const currentDirectory = currentTerminal?.currentDirectory
    ? nonNullArray(currentTerminal?.currentDirectory)
    : undefined;

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.code === "Space" && data?.pageState?.nextStep) {
        router.push(`./?step=${data.pageState.nextStep}`);
      }
    };
    document.addEventListener("keyup", handleKeyDown);

    // Don't forget to clean up
    return function cleanup() {
      document.removeEventListener("keyup", handleKeyDown);
    };
  }, [step, data?.pageState?.nextStep]);

  // Page load optimization:
  useEffect(() => {
    if (data?.pageState?.nextStep) {
      client
        .query({
          query: PageQuery,
          variables: { step: data?.pageState?.nextStep },
        })
        .catch((error) => console.log(error));
    }
  }, [data?.pageState?.nextStep]);

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
            {data?.pageState?.sourceCode && (
              <SourceCodeViewer
                fragment={data.pageState.sourceCode}
                currentDirectory={currentDirectory}
              />
            )}
            {currentTerminal && (
              <TerminalComponent fragment={currentTerminal} />
            )}
            {data.pageState?.nextStep && (
              <button type="button">
                <Link href={`./?step=${data.pageState.nextStep}`}>
                  next step
                </Link>
              </button>
            )}
          </div>
        </main>
      </>
    )
  );
}

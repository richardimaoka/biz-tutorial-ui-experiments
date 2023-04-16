import { useQuery } from "@apollo/client";
import { css } from "@emotion/react";
import Link from "next/link";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import { Header } from "../components/Header";
import { SourceCodeViewer } from "../components/sourcecode/SourceCodeViewer";
import { TerminalComponent } from "../components/terminal/TerminalComponent";
import { graphql } from "../libs/gql";

const PageQuery = graphql(/* GraphQL */ `
  query PageQuery($step: String) {
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

export default function Home() {
  // const [openFilePath, setOpenFilePath] = useState<string>("");
  const router = useRouter();
  const { step } = router.query;
  const stepVariable = typeof step === "string" ? step : undefined;

  const { loading, error, data, client } = useQuery(PageQuery, {
    variables: { step: stepVariable },
  });

  const currentPage = data?.pageState;
  const nextStep = currentPage?.nextStep;
  const prevStep = currentPage?.prevStep;

  const terminals = currentPage?.terminals;
  const [currentTerminalIndex] = useState(0);
  const currentTerminal = terminals && terminals[currentTerminalIndex];
  const currentDirectory = currentTerminal?.currentDirectory
    ? currentTerminal.currentDirectory
    : undefined;

  // console.log("rendering home", prevStep, step, nextStep);

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.code === "Space") {
        if (event.shiftKey) {
          prevStep && router.push(`./?step=${prevStep}&skipAnimation=true`);
        } else {
          nextStep && router.push(`./?step=${nextStep}`);
        }
      }
    };
    document.addEventListener("keyup", handleKeyDown);

    // Don't forget to clean up
    return function cleanup() {
      document.removeEventListener("keyup", handleKeyDown);
    };
  }, [router, step, nextStep, prevStep]);

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
    if (prevStep) {
      client
        .query({
          query: PageQuery,
          variables: { step: prevStep },
        })
        .catch((error) => console.log(error));
    }
  }, [client, nextStep, prevStep]);

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
              /* margin: 0 auto; */
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

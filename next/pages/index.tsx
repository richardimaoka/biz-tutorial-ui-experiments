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
  query PageQuery($step: Int!) {
    terminal(step: $step) {
      ...TerminalComponent_Fragment
    }
    step(stepNum: $step) {
      sourceCode {
        ...SourceCodeViewer_Fragment
      }
      terminals {
        name
        currentDirectory
      }
    }
  }
`);

export default function Home() {
  const router = useRouter();
  const { step } = router.query;
  const stepInt = typeof step === "string" ? Math.trunc(Number(step)) : 1;
  const [terminalName] = useState("default");

  const { loading, error, data } = useQuery(PageQuery, {
    variables: { step: stepInt },
  });

  const terminal = data?.step?.terminals?.find((e) => e?.name === terminalName);
  const currentDirectory = terminal?.currentDirectory
    ? nonNullArray(terminal.currentDirectory)
    : undefined;

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.code === "Space") {
        router.push(`./?step=${stepInt + 1}`);
      }
    };
    document.addEventListener("keydown", handleKeyDown);

    // Don't forget to clean up
    return function cleanup() {
      document.removeEventListener("keydown", handleKeyDown);
    };
  }, [step]);

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
            {data?.step?.sourceCode && (
              <SourceCodeViewer
                fragment={data.step.sourceCode}
                currentDirectory={currentDirectory}
              />
            )}
            {data.terminal && <TerminalComponent fragment={data.terminal} />}
            <button type="button">
              <Link href={`./?step=${stepInt + 1}`}> next step</Link>
            </button>
          </div>
        </main>
      </>
    )
  );
}

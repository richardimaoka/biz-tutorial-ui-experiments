import { useQuery } from "@apollo/client";
import { css } from "@emotion/react";
import { useRouter } from "next/router";
import { Header } from "../components/Header";
import { SourceCodeViewer } from "../components/sourcecode/SourceCodeViewer";
import { graphql } from "../libs/gql";

const PageQuery = graphql(/* GraphQL */ `
  query PageQuery($step: Int!) {
    step(stepNum: $step) {
      sourceCode {
        ...SourceCodeViewer_Fragment
      }
    }
  }
`);

export default function Home() {
  const router = useRouter();
  const { step } = router.query;
  const stepInt = typeof step === "string" ? Math.trunc(Number(step)) : 1;

  const { loading, error, data } = useQuery(PageQuery, {
    variables: { step: stepInt },
  });
  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error : {error.message}</p>;

  return (
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
            <SourceCodeViewer fragment={data.step.sourceCode} />
          )}
        </div>
      </main>
    </>
  );
}

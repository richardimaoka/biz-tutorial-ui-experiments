import { GqlTutorialComponent } from "@/app/components/tutorial/GqlTutorialComponent";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTestTutorialColumnsPage {
    _test {
      appTestTutorialColumnsPage {
        ...GqlTutorialComponent
      }
    }
  }
`);

interface PageParams {
  searchParams: {};
}

export default async function Page({ searchParams }: PageParams) {
  const gqlEndPoint = process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT;
  if (typeof gqlEndPoint != "string") {
    throw new Error("Next.js server gave wrong GraphQL endpoint URL.");
  }

  // const variables = { step: stepNum };
  const data = await request(
    gqlEndPoint,
    queryDefinition
    // variables
  );

  const fragment = data._test?.appTestTutorialColumnsPage;

  return (
    <div style={{ height: "100svh" }}>
      {fragment && <GqlTutorialComponent fragment={fragment} />}
    </div>
  );
}

import { GqlTutorialComponent } from "@/app/components/tutorial/GqlTutorialComponent";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTutorialPage($tutorial: String!) {
    page(tutorial: $tutorial) {
      __typename
      ...GqlTutorialComponent
    }
  }
`);

/**
 * page.js
 * https://nextjs.org/docs/app/api-reference/file-conventions/page
 *
 *
 */
interface PageParams {
  params: {
    tutorial: string;
  };
  searchParams: {
    step?: string;
  };
}

export default async function Page(props: PageParams) {
  const gqlEndPoint = process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT;
  if (typeof gqlEndPoint != "string") {
    throw new Error("Next.js server gave wrong GraphQL endpoint URL.");
  }

  // const variables = { step: stepNum };
  const data = await request(gqlEndPoint, queryDefinition, {
    tutorial: props.params.tutorial,
  });

  const fragment = data.page;

  return (
    <div style={{ height: "100svh" }}>
      {fragment && <GqlTutorialComponent fragment={fragment} />}
    </div>
  );
}

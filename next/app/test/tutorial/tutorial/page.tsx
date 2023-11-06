import { GqlTutorialComponent } from "@/app/components/tutorial/GqlTutorialComponent";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTestTutorialTutorialPage {
    _test {
      appTestTutorialTutorialPage {
        ...GqlTutorialComponent
      }
    }
  }
`);

interface PageParams {
  searchParams: {};
}

export default async function Page({ searchParams }: PageParams) {
  // const variables = { step: stepNum };
  const data = await request(
    "http://localhost:8080/query",
    queryDefinition
    // variables
  );

  const fragment = data._test?.appTestTutorialTutorialPage;

  return <div>{fragment && <GqlTutorialComponent fragment={fragment} />}</div>;
}

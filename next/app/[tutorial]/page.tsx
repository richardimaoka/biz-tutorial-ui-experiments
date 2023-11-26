import { GqlTutorialComponent } from "@/app/components/tutorial/GqlTutorialComponent";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";
import { HandleKeyEvents } from "../components/navigation/HandleKeyEvents";
import { print } from "graphql";
import { HandleTrivial } from "../components/navigation/HandleTrivial";
import styles from "./page.module.css";

const queryDefinition = graphql(`
  query appTutorialPage($tutorial: String!, $step: String) {
    page(tutorial: $tutorial, step: $step) {
      __typename
      nextStep
      prevStep
      isTrivial
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
  // console.log("----------------------------------------");
  // console.log(print(queryDefinition));

  const gqlEndPoint = process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT;
  if (typeof gqlEndPoint != "string") {
    throw new Error("Next.js server gave wrong GraphQL endpoint URL.");
  }

  const variables = {
    tutorial: props.params.tutorial,
    step: props.searchParams.step,
  };
  const data = await request(gqlEndPoint, queryDefinition, variables);

  const fragment = data.page;

  return (
    <div className={styles.component}>
      {fragment && <GqlTutorialComponent fragment={fragment} />}
      <HandleKeyEvents
        tutorial={props.params.tutorial}
        prevStep={fragment?.prevStep}
        nextStep={fragment?.nextStep}
      />
      <HandleTrivial
        isTrivial={fragment?.isTrivial}
        nextStep={fragment?.nextStep}
      />
    </div>
  );
}

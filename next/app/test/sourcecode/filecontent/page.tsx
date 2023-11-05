import { Carousel } from "@/app/components/carousel/Carousel";
import { GqlOpenFilePane } from "@/app/components/sourcecode2/openfile/GqlOpenFilePane";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";
import Link from "next/link";

const queryDefinition = graphql(`
  query appTestSourcecodeFilecontentPage($step: Int!) {
    _test {
      appTestSourcecodeFilecontentPage(step: $step) {
        ...GqlOpenFilePane
      }
    }
  }
`);

interface PageParams {
  searchParams: {
    step?: string;
  };
}

export default async function Page({ searchParams }: PageParams) {
  const stepNum = searchParams.step ? Number(searchParams.step) : 1;
  const nextStep = `${stepNum + 1}`;

  const variables = { step: stepNum };
  const data = await request(
    "http://localhost:8080/query",
    queryDefinition,
    variables
  );

  const fragment = data._test?.appTestSourcecodeFilecontentPage;

  if (!fragment) {
    return <>no data</>;
  }

  return (
    <div style={{ height: "95svh" }}>
      <Link href={`/test/sourcecode/filecontent?step=${nextStep}`}>next</Link>
      <Carousel currentIndex={0} columnWidth={768}>
        <div style={{ display: "flex", height: "100%" }}>
          <GqlOpenFilePane fragment={fragment} />
          <GqlOpenFilePane fragment={fragment} />
        </div>
      </Carousel>
    </div>
  );
}

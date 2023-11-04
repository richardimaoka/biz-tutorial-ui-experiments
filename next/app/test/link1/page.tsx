import Link from "next/link";
import { MyLink } from "../../components/link/MyLink";
import { NextButton } from "@/app/components/navigation2/NextButton";

interface PageParams {
  searchParams: {
    p: string;
  };
}

export default async function Page({ searchParams }: PageParams) {
  const pNumber = searchParams.p ? Number(searchParams.p) : 0;
  return (
    <>
      <MyLink path="/test/link2" searchParams={{ p: `${pNumber + 1}` }}>
        aaa
      </MyLink>
      <Link href={`/test/link2`}>go</Link>
      <NextButton href={`/test/link2?p=${pNumber + 1}`} />
    </>
  );
}

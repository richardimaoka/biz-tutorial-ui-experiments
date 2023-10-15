import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

// This function can be marked `async` if using `await` inside
export function middleware(request: NextRequest) {
  if (request.nextUrl.pathname.startsWith("/test")) {
    if (process.env.NODE_ENV === "production") {
      return NextResponse.json(
        {
          success: false,
          message: "/test path cannot be executed in production",
        },
        { status: 400 }
      );
    }
  }
}

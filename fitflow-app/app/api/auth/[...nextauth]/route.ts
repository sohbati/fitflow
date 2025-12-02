// NextAuth API route removed - using custom IAM service instead
// This file is kept for potential future use or reference

import { NextResponse } from 'next/server';

// Return 404 since this route is not implemented
export async function GET() {
  return NextResponse.json(
    { error: 'This route is not implemented. Using custom IAM service instead.' },
    { status: 404 }
  );
}

export async function POST() {
  return NextResponse.json(
    { error: 'This route is not implemented. Using custom IAM service instead.' },
    { status: 404 }
  );
}

"use client";

import { useMutation, useQuery } from "@tanstack/react-query";
import Image from "next/image";
import { useEffect } from "react";
import { Button } from "~/components/ui/button";
import { CardContent, Card } from "~/components/ui/card";

type Product = {
  id: number;
  name: string;
  price: number;
};

const getProducts = async () => {
  const response = await fetch("http://localhost/api/merchant/products");
  const result: { products: Product[] } = await response.json();

  return result.products;
};

const purchase = async (product: Product) => {
  const response = await fetch("http://localhost/api/checkout/purchase", {
    method: "POST",
    body: JSON.stringify({
      address: "dsadas",
      email: "dsadsa@gmail.com",
      products: [product],
    }),
    headers: {
      "Content-Type": "application/json",
    }
  });

  const result = await response.json();

  return result;
};

export default function Component() {
  const { data, isLoading } = useQuery({
    queryKey: ["products"],
    queryFn: getProducts,
  });

  const { mutate: handleCheckout } = useMutation({
    mutationKey: ["checkout"],
    mutationFn: purchase,
    onSuccess: () => {
      alert("Checkout success");
    },
    onError: () => {
      alert("Checkout failed");
    },
  });

  if (isLoading) return <div>Loading...</div>;

  return (
    <div className="container mx-auto">
      <div className="py-14">
        <h1 className="text-4xl font-bold mb-4">Checkout Journey</h1>
        <p className="w-1/2 text-zinc-300">
          This app is an experimental checkout journey simulation. It uses some
          technologies to achieve seamless checkout experience. These products
          are fetched from the <b>merchant</b> API.
          <br />
          <i>Educational purpose only. </i>
          <br /> <br />
          When you checkout, the <b>checkout</b> API will be called to emit the{" "}
          <code>checkout</code> event. Following service will consume it:
        </p>

        <ul className="list-disc list-inside mt-4">
          <li>Notification: Send email to the customer</li>
          <li>Inventory: Reduce the stock of the product</li>
          <li>Shipment: Create a shipment for the order</li>
        </ul>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 p-4">
        {data &&
          data.map((product) => (
            <Card key={product.id}>
              <CardContent>
                <Image
                  alt={product.name}
                  className="w-full h-64 object-cover my-4 rounded-md"
                  height="200"
                  src={`https://placehold.co/600x600/png?text=${product.name}`}
                  style={{
                    aspectRatio: "200/200",
                    objectFit: "cover",
                  }}
                  width="200"
                />
                <h2 className="font-semibold text-lg mb-2">{product.name}</h2>
                <p className="text-gray-500 mb-2">{product.id}</p>
                <p className="text-xl mb-4">${product.price}</p>
                <Button onClick={() => handleCheckout(product)}>
                  Checkout
                </Button>
              </CardContent>
            </Card>
          ))}
      </div>
    </div>
  );
}

import { BuyProductService, ListProductService } from '@/services/product';
import { IListItems } from '@/store/typings/root';
import { useEffect, useState } from 'react';

interface ICart {
  id: number;
  image_url: string;
  quantity: number;
}
const Home = () => {
  const [list, setList] = useState<IListItems[]>([]);
  const [cart, setCart] = useState<ICart[]>([]);
  const [money, setMoney] = useState({
    value: [
      { type: 1, count: 0 },
      { type: 5, count: 0 },
      { type: 10, count: 0 },
      { type: 20, count: 0 },
      { type: 50, count: 0 },
      { type: 100, count: 0 },
      { type: 500, count: 0 },
      { type: 1000, count: 0 }
    ],
    total: 0
  });

  const fetchData = async () => {
    try {
      const { data } = await ListProductService();
      setList(data.results);
    } catch (error) {}
  };
  useEffect(() => {
    fetchData();
  }, []);

  const handleAddToCart = (id: number) => {
    const index = cart.findIndex((item) => item.id === id);
    if (index === -1) {
      setCart([
        ...cart,
        { id, quantity: 1, image_url: list.find((item) => item.id === id)?.image_url || '' }
      ]);
    } else {
      const newCart = cart.map((item) => {
        if (item.id === id) {
          return { ...item, quantity: item.quantity + 1 };
        }
        return item;
      });
      setCart(newCart);
    }
  };

  const handleAddQuantity = (id: number, type: 'add' | 'remove') => {
    const item = cart.find((item) => item.id === id);

    if (type === 'remove' && item?.quantity === 1) {
      const newCart = cart.filter((item) => item.id !== id);
      setCart(newCart);
    } else {
      const newCart = cart.map((item) => {
        if (item.id === id) {
          return { ...item, quantity: type === 'add' ? item.quantity + 1 : item.quantity - 1 };
        }
        return item;
      });
      setCart(newCart);
    }
  };

  const handleAddMoney = (value: number) => {
    const newMoney = money.value.map((item) => {
      if (item.type === value) {
        return { ...item, count: item.count + 1 };
      }
      return item;
    });
    const total = newMoney.reduce((acc, item) => acc + item.type * item.count, 0);
    setMoney({ value: newMoney, total });
  };

  const handleCheckOut = async () => {
    const total = cart.reduce((acc, item) => {
      const product = list.find((product) => product.id === item.id);
      return acc + (product?.price || 0) * item.quantity;
    }, 0);

    const newMoney = money.value.filter((item) => item.count > 0);
    const newCart = cart.map((item) => ({
      id: item.id,
      quantity: item.quantity
    }));

    const pay = money.value.reduce((acc, item) => acc + item.type * item.count, 0);
    const payload = {
      product: newCart,
      value: newMoney,
      pay,
      total
    };
    if (total > money.total) {
      alert('Not enough money');
    } else {
      try {
        await BuyProductService(payload);
        alert('Check out success');
      } catch (error) {
        alert('Check out failed');
      }
    }
  };

  return (
    <div className="flex items-center justify-center w-full h-screen">
      <div className="grid w-3/6 grid-cols-12 overflow-y-auto h-5/6">
        {list.map((item) => (
          <div
            key={item.id}
            className="flex flex-col justify-between col-span-12 m-1 border md:col-span-6 lg:col-span-4 rounded-xl border-neutral-300">
            <img src={item.image_url} alt={item.name} className="object-cover w-full h-40 mt-2" />
            <div>
              <p className="mt-2 text-sm text-center">{item.name}</p>
              <p className="mt-2 text-sm text-center">
                {item.price} BAHT | {item.stock} in stock{' '}
              </p>
              <div className="flex justify-center">
                <p
                  className="w-3/6 p-2 my-2 text-xs text-center text-white bg-red-400 rounded-lg cursor-pointer hover:bg-red-200"
                  onClick={() => handleAddToCart(item.id)}>
                  Add to cart
                </p>
              </div>
            </div>
          </div>
        ))}
      </div>

      <div className="flex flex-col justify-between w-1/6 p-2 border rounded-lg border-neutral-300 h-5/6">
        <div>
          <div className="mb-2">
            <div>
              <span className="text-xs">Add Money</span>
              <span>
                <span className="p-1 m-1 text-xs bg-green-400 rounded-lg cursor-pointer">
                  {money.total}
                </span>
              </span>
            </div>
            <div className="flex flex-col">
              <div className="flex flex-wrap">
                <span
                  className="p-1 m-1 text-xs bg-gray-400 rounded-lg cursor-pointer"
                  onClick={() => handleAddMoney(1)}>
                  + 1
                </span>
                <span
                  className="p-1 m-1 text-xs bg-gray-400 rounded-lg cursor-pointer"
                  onClick={() => handleAddMoney(5)}>
                  + 5
                </span>
                <span
                  className="p-1 m-1 text-xs bg-gray-400 rounded-lg cursor-pointer"
                  onClick={() => handleAddMoney(10)}>
                  + 10
                </span>
                <span
                  className="p-1 m-1 text-xs bg-gray-400 rounded-lg cursor-pointer"
                  onClick={() => handleAddMoney(20)}>
                  + 20
                </span>
                <span
                  className="p-1 m-1 text-xs bg-gray-400 rounded-lg cursor-pointer"
                  onClick={() => handleAddMoney(50)}>
                  + 50
                </span>
                <span
                  className="p-1 m-1 text-xs bg-gray-400 rounded-lg cursor-pointer"
                  onClick={() => handleAddMoney(100)}>
                  + 100
                </span>

                <span
                  className="p-1 m-1 text-xs bg-gray-400 rounded-lg cursor-pointer"
                  onClick={() => handleAddMoney(500)}>
                  + 500
                </span>

                <span
                  className="p-1 m-1 text-xs bg-gray-400 rounded-lg cursor-pointer"
                  onClick={() => handleAddMoney(1000)}>
                  + 1000
                </span>
              </div>
            </div>
          </div>
          <div className="flex justify-between">
            <div>
              <span>
                <span className="text-sm">Cart </span>
                {cart.length > 0 && <span className="text-sm">({cart.length}) items</span>}
              </span>
            </div>
            <div>
              <span className="text-sm">
                Total{' '}
                {cart.reduce((acc, item) => {
                  const product = list.find((product) => product.id === item.id);
                  return acc + (product?.price || 0) * item.quantity;
                }, 0)}
              </span>
            </div>
          </div>
          <div className="mt-2 overflow-y-auto max-h-80">
            {cart.map((item) => {
              const product = list.find((product) => product.id === item.id);
              return (
                <div key={item.id} className="flex items-center justify-between p-2">
                  <img
                    src={item.image_url}
                    alt={product?.name}
                    className="object-cover w-20 h-20"
                  />
                  <div>
                    <span
                      className="px-2 py-1 text-white bg-red-500 rounded-l-lg cursor-pointer"
                      onClick={() => handleAddQuantity(item.id, 'remove')}>
                      -
                    </span>
                    <span className="px-2 py-1 border border-y-2">{item.quantity}</span>
                    {product?.stock && product.stock > item.quantity && (
                      <span
                        className="px-2 py-1 text-white bg-green-500 rounded-r-lg cursor-pointer"
                        onClick={() => handleAddQuantity(item.id, 'add')}>
                        +
                      </span>
                    )}
                  </div>
                </div>
              );
            })}
          </div>
        </div>

        <div className="flex justify-center w-full uppercase">
          <span
            className="flex justify-center w-4/6 p-2 text-white uppercase cursor-pointer md:w-5/6 bg-neutral-400 hover:bg-neutral-300 rounded-3xl"
            onClick={handleCheckOut}>
            check out
          </span>
        </div>
      </div>
    </div>
  );
};

export default Home;

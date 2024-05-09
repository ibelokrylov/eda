import puppeteer from 'puppeteer';
import fs from 'fs';
import path from 'path';

const site_product = 'https://calorizator.ru/product';

const randomIntFromInterval = (min, max) => {
  return Math.floor(Math.random() * (max - min + 1) + min);
};

const createFile = async (data, name) => {
  const dateFormat = new Date().toISOString().split('T')[0];

  const pathFolder = path.join(process.cwd(), 'scraper', 'data', dateFormat);

  if (!fs.existsSync(pathFolder)) await fs.mkdirSync(pathFolder);

  const filePath = path.join(process.cwd(), 'scraper', 'data', dateFormat, `${dateFormat}-${name ?? 'products'}.json`);
  await fs.writeFileSync(filePath, JSON.stringify(data));
};

const createBrowser = async () => {
  const browser = await puppeteer.launch({
    headless: true,
  });

  return {
    page: await browser.newPage(),
    browser,
  };
};

const getDataProduct = async (page) => {
  return await page.evaluate((i) => {
    const body = document.body;
    const product_name = body.querySelectorAll('td.views-field-title a');
    const _pr_name = Array.from(product_name).map((el) => el.innerHTML);
    const protein = body.querySelectorAll('td.views-field-field-protein-value');
    const fat = body.querySelectorAll('td.views-field-field-fat-value');
    const carbohydrate = body.querySelectorAll('td.views-field-field-carbohydrate-value');
    const kcal = body.querySelectorAll('td.views-field-field-kcal-value');

    return _pr_name.map((item, index) => {
      return {
        name: item,
        protein: protein[index].innerText,
        fat: fat[index].innerText,
        carbohydrate: carbohydrate[index].innerText,
        kcal: kcal[index].innerText,
      };
    });
  });
};

const getProducts = async (url, page) => {
  console.log('🚀 ~ getProducts ~ url:', url);
  await page.goto(url, {
    waitUntil: 'domcontentloaded',
  });

  const pages = await page.evaluate(() => {
    const p = document.body.querySelectorAll('.pager li').length;
    return document.body.querySelectorAll('.pager li').length === 0 ? 0 : p - 2;
  });
  console.log('🚀 ~ pages ~ pages:', pages);
  const data = [];

  if (pages) {
    for (let i = 0; i <= pages; i++) {
      if (i !== 0) {
        const _url = `${url}?page=${i}`;
        await page.waitForTimeout(randomIntFromInterval(3000, 8000));
        await page.goto(_url, {
          waitUntil: 'domcontentloaded',
        });
      } else {
        await page.waitForTimeout(randomIntFromInterval(3000, 8000));
      }
      const pr = await getDataProduct(page);
      setTimeout(() => {
        data.push(...pr);
      }, 0);
    }
  } else {
    const pr = await getDataProduct(page);
    setTimeout(() => {
      data.push(...pr);
    }, 0);
  }
  return data;
};

const categoriesAndProduct = async () => {
  const { page, browser } = await createBrowser();
  await page.goto(site_product, {
    waitUntil: 'domcontentloaded',
  });
  const categories = await page.evaluate(() => {
    const products = document.querySelectorAll('.product li');
    return Array.from(products).map((item) => {
      return {
        title: item.querySelector('a').innerHTML,
        link: item.querySelector('a').href,
      };
    });
  });
  categories.splice(categories.length - 5, 5);
  const categoriesLength = categories.length;
  for (let i = 0; i < categoriesLength; i++) {
    const products = await getProducts(categories[i].link, page);
    setTimeout(() => {
      createFile(
        {
          title: categories[i].title,
          products,
        },
        categories[i].title
      );
    }, 0);
  }
  await browser.close();
};
// console.log(await categoriesAndProduct());
console.log(new Date().toISOString().split('T')[0]);
